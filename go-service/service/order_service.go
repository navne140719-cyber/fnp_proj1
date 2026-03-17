package service

import (
	"errors"
	"fmt"
	"go-service/config"
	"go-service/kafka"
	"go-service/models"
	"go-service/repository"
)

func CreateOrder(req models.OrderRequest) (int64, error) {

	if req.UserID <= 0 {
		return 0, errors.New("invalid user id")
	}

	if len(req.Items) == 0 {
		return 0, errors.New("order must contain at least one item")
	}

	// MERGE DUPLICATE PRODUCTS

	productQtyMap := make(map[int]int)

	for _, item := range req.Items {
		productQtyMap[item.ProductID] += item.Qty
	}

	//  VALIDATE + CALCULATE TOTAL

	var total float64

	for productID, qty := range productQtyMap {

		// invalid quantity
		if qty <= 0 {
			return 0, fmt.Errorf("invalid quantity for product %d", productID)
		}

		// unreal quantity
		if qty > 1000 {
			return 0, fmt.Errorf("quantity too large for product %d", productID)
		}

		//  fetch price
		row := config.DB.QueryRow(
			"SELECT price, stock FROM products WHERE id=?",
			productID,
		)

		var price float64
		var stock int

		err := row.Scan(&price, &stock)
		if err != nil {
			return 0, fmt.Errorf("product not found: %d", productID)
		}

		//  stock validation
		if stock < qty {
			return 0, fmt.Errorf("insufficient stock for product %d", productID)
		}

		total += price * float64(qty)
	}

	// CREATE ORDER

	orderID, err := repository.CreateOrder(req.UserID, total)
	if err != nil {
		return 0, err
	}

	//  INSERT ORDER ITEMS

	for productID, qty := range productQtyMap {

		row := config.DB.QueryRow(
			"SELECT price FROM products WHERE id=?",
			productID,
		)

		var price float64
		err := row.Scan(&price)
		if err != nil {
			return 0, err
		}

		err = repository.InsertOrderItem(orderID, productID, qty, price)
		if err != nil {
			return 0, err
		}
	}

	// CREATE KAFKA EVENT

	event := map[string]interface{}{
		"orderId": orderID,
		"userId":  req.UserID,
		"items":   req.Items,
	}

	// PUBLISH EVENT

	err = kafka.PublishOrderEvent(event)
	if err != nil {
		return 0, err
	}

	return orderID, nil
}
