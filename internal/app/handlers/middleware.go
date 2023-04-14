package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/srjchsv/chat-service/internal/app/services"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			return
		}

		userID, username, err := services.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Set("user_id", userID)
		c.Set("username", username)
		c.Next()
	}
}
