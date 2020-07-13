package middlewares

import (
	"net/http"

	"go-api/pkg"
	"github.com/gin-gonic/gin"
)

type HttpHandlerFunc func(*gin.Context)

func Auth(c *gin.Context) {
        token := c.GetHeader("Token")
	if token == pkg.GetEnvOrDefault("API_TOKEN", "hunter2") {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
