package service

import (
	"go-service/config"
	"go-service/kafka"
	"go-service/models"
	"go-service/repository"
)

func CreateOrder(req models.OrderRequest) (int64, error) {
	var total float64
	for _, item := range req.Items {
		row := config.DB.QueryRow(
			"SELECT price FROM products WHERE id=?",
			item.ProductID,
		)
		var price float64
		err := row.Scan(&price)
		if err != nil {
			return 0, err
		}
		total += price * float64(item.Qty)
	}

	// Create order
	orderID, err := repository.CreateOrder(req.UserID, total)
	if err != nil {
		return 0, err
	}
	// Insert order items
	for _, item := range req.Items {
		row := config.DB.QueryRow(
			"SELECT price FROM products WHERE id=?",
			item.ProductID,
		)
		var price float64
		err := row.Scan(&price)
		if err != nil {
			return 0, err
		}
		err = repository.InsertOrderItem(orderID, item.ProductID, item.Qty, price)
		if err != nil {
			return 0, err
		}
	}

	// Create Kafka Event
	event := map[string]interface{}{
		"orderId": orderID,
		"userId":  req.UserID,
		"items":   req.Items,
	}

	// Publish event to Kafka
	err = kafka.PublishOrderEvent(event)
	if err != nil {
		return 0, err
	}
	return orderID, nil
}
