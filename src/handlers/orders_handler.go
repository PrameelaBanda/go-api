package handlers

import (
	"encoding/json"
	"fmt"
	"go-api/src/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"io"
	"log"
    "encoding/csv"
)

func HandleListOrders(r *repository.Repository) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		ordersList, err := r.ListOrders()
		if err != nil {
			if dbErr, ok := err.(repository.Dberror); ok {
				c.AbortWithStatusJSON(dbErr.Code, err)
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		b, err := json.MarshalIndent(ordersList, "", " ")
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

func HandlePostData(r *repository.Repository) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		var result [][]string
		var err error
		fileName := c.Param("fileName")
		resp := csv.NewReader(c.Request.Body)
		for {
			record, err := resp.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return
			}
			result = append(result, record)
		}
		switch fileName {
		case "orders" : err = PostOrderData(result, r)
		case "customers" : err = PostCustomerData(result, r)
		case "deliveries" : err = PostDeliveriesData(result, r)
		case "orderItems" : err = PostOrderItemsData(result, r)
		case "customerCompanies" : err = PostCustomerCompanies(result, r)
		}
		if err != nil {
			if dbErr, ok := err.(repository.Dberror); ok {
				c.AbortWithStatusJSON(dbErr.Code, err)
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Data loaded successfully"),
		})
	}

	return fn
}

func DecodeFileToString(result []string) ([]string) {
	var row []string
	b, err := json.Marshal(result)
	if err != nil {
		log.Fatal("\nERROR while marshaling the data ", err)
	}
	err = json.Unmarshal(b, &row)
	if err != nil {
		log.Fatal("\nERROR while unmarshaling the data ", err)
	}
	return row
}
