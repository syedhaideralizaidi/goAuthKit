package controller

import (
	"context"
	"encoding/base64"
	"errors"
	"math/rand"
	"net/http"
	"time"

	db "github.com/syedhaideralizaidi/authkit/internal/database/sqlc"
	"github.com/syedhaideralizaidi/authkit/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

// generateResetToken generates a secure reset token
func (c *Controller) generateResetToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// RequestPasswordReset handles password reset request
func (c *Controller) RequestPasswordReset(ctx *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := c.Queries.GetUserByEmail(context.Background(), req.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	token, err := c.generateResetToken()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	expiry := time.Now().Add(1 * time.Hour)

	requestResetToken := db.UpdateResetTokenParams{
		Email:            req.Email,
		ResetTokenExpiry: pgtype.Timestamptz{Time: expiry, Valid: true},
		ResetToken:       pgtype.Text{String: token, Valid: true},
	}

	err = c.Queries.UpdateResetToken(context.Background(), requestResetToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update reset token"})
		return
	}

	err = utils.SendResetPasswordEmail(req.Email, token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send reset email"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password reset email sent"})
}

// ResetPassword handles password reset
func (c *Controller) ResetPassword(ctx *gin.Context) {
	var req struct {
		Email       string `json:"email" binding:"required,email"`
		Token       string `json:"token" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	requestResetPassword := db.ResetPasswordParams{
		Password:   string(hashedPassword),
		ResetToken: pgtype.Text{String: req.Token, Valid: true},
		Email:      req.Email,
	}

	_, err = c.Queries.ResetPassword(context.Background(), requestResetPassword)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password reset successful"})
}
