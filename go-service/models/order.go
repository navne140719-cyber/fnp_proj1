package models

type OrderItemRequest struct {
	ProductID int `json:"productId"`
	Qty       int `json:"qty"`
}

type OrderRequest struct {
	UserID int                `json:"userId"`
	Items  []OrderItemRequest `json:"items"`
}
