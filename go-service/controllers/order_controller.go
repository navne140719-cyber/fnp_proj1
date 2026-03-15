package controllers

import (
	"net/http"

	"go-service/models"
	"go-service/service"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var req models.OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	orderID, err := service.CreateOrder(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Order Created",
		"orderId": orderID,
	})
}
