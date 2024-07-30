package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/syedhaideralizaidi/authkit/internal/utils"
)

func SetupLogger() *logrus.Logger {
	logger := logrus.New()

	// Create or open a log file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatal(err)
	}

	// Set logrus to log to the file
	logger.SetOutput(file)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetLevel(logrus.InfoLevel) // Set appropriate level

	return logger
}

func AuthMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log the incoming request
		logger.Infof("Incoming request: %s %s", c.Request.Method, c.Request.URL.Path)

		// Your existing middleware logic
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logger.Warn("Authorization header is required")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			logger.Warn("Invalid Authorization header format")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			logger.Warn("Invalid token")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user", claims)
		logger.Infof("Authenticated user: %s", claims.Username)
		c.Next()
	}
}
