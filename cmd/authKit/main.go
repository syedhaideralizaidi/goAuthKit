package main

import (
	"context"

	"github.com/syedhaideralizaidi/authkit/internal/controller"
	"github.com/syedhaideralizaidi/authkit/internal/database"
	"github.com/syedhaideralizaidi/authkit/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := middleware.SetupLogger()
	database.ConnectDB()
	defer database.Conn.Close(context.Background())

	r := gin.Default()
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/verify-email", controller.VerifyEmail)
	r.POST("/request-password-reset", controller.RequestPasswordReset)
	r.POST("/reset-password", controller.ResetPassword)

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware(logger))
	{
		protected.POST("/users", controller.CreateUser)
		protected.GET("/users/:id", controller.GetUser)
		protected.PUT("/users/:id", controller.UpdateUser)
		protected.DELETE("/users/:id", controller.DeleteUser)
		protected.GET("/users", controller.ListUsers)
	}

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
