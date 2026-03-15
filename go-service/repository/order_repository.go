package repository

import "go-service/config"

func CreateOrder(userID int, total float64) (int64, error) {
	result, err := config.DB.Exec(
		"INSERT INTO orders(user_id,status,total_amount) VALUES(?,?,?)",
		userID,
		"CREATED",
		total,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func InsertOrderItem(orderID int64, productID int, qty int, price float64) error {
	_, err := config.DB.Exec(
		"INSERT INTO order_items(order_id,product_id,qty,price) VALUES(?,?,?,?)",
		orderID,
		productID,
		qty,
		price,
	)
	return err
}
