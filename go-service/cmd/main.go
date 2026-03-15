package main

import (
	"go-service/config"
	"go-service/controllers"
	"go-service/kafka"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	kafka.InitProducer()
	router := gin.Default()
	router.GET("/products", controllers.GetProducts)
	router.POST("/orders", controllers.CreateOrder)
	router.Run(":8080")
}
