package route

import (
	"github.com/gin-gonic/gin"
	"github.com/syedhaideralizaidi/authkit/internal/controller"
	"github.com/syedhaideralizaidi/authkit/internal/database"
	"github.com/syedhaideralizaidi/authkit/internal/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	userController := controller.NewController(database.Queries)
	logger := middleware.SetupLogger()
	SetupPublicRoutes(router, userController)

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware(logger))
	SetupAuthenticatedRoutes(protected, userController)
	return router
}
