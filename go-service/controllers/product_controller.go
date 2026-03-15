package controllers

import (
	"net/http"

	"go-service/service"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products, err := service.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, products)
}
