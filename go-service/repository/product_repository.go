package repository

import (
	"go-service/config"
	"go-service/models"
)

func GetAllProducts() ([]models.Product, error) {
	rows, err := config.DB.Query("SELECT id,name,price,stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []models.Product
	for rows.Next() {
		var p models.Product
		rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
		products = append(products, p)
	}
	return products, nil
}
