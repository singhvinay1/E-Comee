package middleware

import (
	"net/http"
	"strings"

	"ecommerce-backend/database"
	"ecommerce-backend/models"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates the user token and sets the user in the context
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Check if the header starts with "Bearer "
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		token := tokenParts[1]
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}

		// Find user by token
		var user models.User
		if err := database.DB.Where("token = ?", token).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set user in context
		c.Set("user", user)
		c.Next()
	}
}

// GetUserFromContext extracts the user from the gin context
func GetUserFromContext(c *gin.Context) (*models.User, bool) {
	userInterface, exists := c.Get("user")
	if !exists {
		return nil, false
	}

	user, ok := userInterface.(models.User)
	if !ok {
		return nil, false
	}

	return &user, true
} 