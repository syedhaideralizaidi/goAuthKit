package route

import (
	"github.com/gin-gonic/gin"
	"github.com/syedhaideralizaidi/authkit/internal/controller"
)

// SetupAuthenticatedRoutes sets up the routes that require authentication
func SetupAuthenticatedRoutes(rg *gin.RouterGroup, userController *controller.Controller) {
	{
		rg.POST("/users", userController.CreateUser)
		rg.GET("/users/:id", userController.GetUser)
		rg.PUT("/users/:id", userController.UpdateUser)
		rg.DELETE("/users/:id", userController.DeleteUser)
		rg.GET("/users", userController.ListUsers)
	}
}
