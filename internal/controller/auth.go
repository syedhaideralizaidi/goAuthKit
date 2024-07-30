package controller

import (
	"context"
	"net/http"

	db "github.com/syedhaideralizaidi/authkit/internal/database/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"

	"github.com/syedhaideralizaidi/authkit/internal/database"
	"github.com/syedhaideralizaidi/authkit/internal/utils"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email       string `json:"email" binding:"required,email"`
	Username    string `json:"username" binding:"required"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password" binding:"required"`
	Role        string `json:"role" binding:"required,oneof=admin seller buyer"`
}

func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	createUserParams := db.CreateUserParams{
		Email:       req.Email,
		Username:    req.Username,
		PhoneNumber: pgtype.Text{String: req.PhoneNumber, Valid: req.PhoneNumber != ""},
		Password:    string(hashedPassword),
		Role:        req.Role,
		IsVerified:  pgtype.Bool{Bool: false, Valid: true},
	}

	user, err := database.Queries.CreateUser(context.Background(), createUserParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateJWT(user.Username, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	err = utils.SendVerificationEmail(user.Email, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send verification email"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created, verification email sent"})
}

func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := database.Queries.GetUserByEmail(context.Background(), req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect email or password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect email or password"})
		return
	}

	if !user.IsVerified.Bool {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not verified"})
		return
	}

	token, err := utils.GenerateJWT(user.Username, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func VerifyEmail(c *gin.Context) {
	tokenString := c.Query("token")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	claims, err := utils.ValidateJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	err = database.Queries.VerifyUser(context.Background(), claims.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}
