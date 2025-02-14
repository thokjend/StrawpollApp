package middleware

import (
	"net/http"
	"strawpoll-app/database"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates the session token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing session token"})
			c.Abort()
			return
		}

		// Check if the token exists in Redis
		user, err := database.Client.Get(database.Ctx, "session:"+token).Result()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired session token"})
			c.Abort()
			return
		}

		// Store username in context for further use
		c.Set("username", user)
		c.Next()
	}
}