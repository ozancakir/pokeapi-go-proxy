package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func ApiKeyAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")

		if apiKey == "" {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		if apiKey != os.Getenv("API_KEY") {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
