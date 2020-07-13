package handlers

import (
	"encoding/json"
	"go-api/src/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleListOrders(r *repository.Repository) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		podList, err := r.ListOrders()
		if err != nil {
			if dbErr, ok := err.(repository.Dberror); ok {
				c.AbortWithStatusJSON(dbErr.Code, err)
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		b, err := json.MarshalIndent(podList, "", " ")
		if err != nil {
			if dbErr, ok := err.(repository.Dberror); ok {
				c.AbortWithStatusJSON(dbErr.Code, err)
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
		c.Header("Content-Type", "application/json")
		c.String(http.StatusOK, string(b))
	}

	return fn
}

func HandleGetOrder(r *repository.Repository) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		order, err := r.GetOrder(c.Param("orderName"))
		if err != nil {
			if dbErr, ok := err.(repository.Dberror); ok {
				c.AbortWithStatusJSON(dbErr.Code, err)
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		b, err := json.MarshalIndent(order, "", " ")
		if err != nil {
			if dbErr, ok := err.(repository.Dberror); ok {
				c.AbortWithStatusJSON(dbErr.Code, err)
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
		c.Header("Content-Type", "application/json")
		c.String(http.StatusOK, string(b))
	}

	return fn
}