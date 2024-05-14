package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApplyJSONMiddleware(c *gin.Context) {
	if c.ContentType() != "application/json" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request must be JSON"})
		c.Abort()
		return
	}
	c.Next()
}
