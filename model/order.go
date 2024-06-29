package model

import "time"

type Order struct {
	Id       int64     `json:"id" db:"id"`
	Datetime time.Time `json:"datetime" db:"datetime"`
}

type OrderItem struct {
	Id        int64 `json:"id" db:"id"`
	OrderId   int64 `json:"order_id" binding:"required" db:"order_id"`
	ProductId int64 `json:"product" binding:"required" db:"product_id"`
	Quantity  int64 `json:"quantity" binding:"required" db:"quantity"`
}
