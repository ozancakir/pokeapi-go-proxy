package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func ApiKeyAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		apiKey := os.Getenv("API_KEY")
		xApiKey := c.GetHeader("X-API-KEY")

		if apiKey != "" && xApiKey != apiKey {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
