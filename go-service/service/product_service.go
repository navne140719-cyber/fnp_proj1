package service

import (
	"go-service/models"
	"go-service/repository"
)

func GetProducts() ([]models.Product, error) {
	return repository.GetAllProducts()
}
