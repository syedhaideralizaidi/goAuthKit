package route

import (
	"github.com/gin-gonic/gin"
	"github.com/syedhaideralizaidi/authkit/internal/controller"
)

// SetupPublicRoutes sets up the routes that do not require authentication
func SetupPublicRoutes(r *gin.Engine, userController *controller.Controller) {
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)
	r.GET("/verify-email", userController.VerifyEmail)
	r.POST("/request-password-reset", userController.RequestPasswordReset)
	r.POST("/reset-password", userController.ResetPassword)
}
