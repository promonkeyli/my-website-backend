package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"web_backend.com/m/v2/utils"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// Check if the token is provided and starts with "Bearer "
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing or malformed token"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix from the token
		tokenString = tokenString[7:]

		// Parse the token and get the claims
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// Store the username in the context for later use
		c.Set("username", claims.Username)
		c.Next()
	}
}
