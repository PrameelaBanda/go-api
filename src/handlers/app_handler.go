package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleApp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"apiVersion": "v1",
		"resources":  []string{"pod"},
		"ops":        []string{"GET", "POST", "PUT", "DELETE"},
	})

}
