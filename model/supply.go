package model

import "time"

type Supply struct {
	Id       int64     `json:"id" db:"id"`
	Datetime time.Time `json:"timestamp" db:"timestamp"`
}

type SupplyItem struct {
	Id        int64 `json:"id" db:"id"`
	SupplyId  int64 `json:"supply_id" binding:"required" db:"supply_id"`
	ProductId int64 `json:"product" binding:"required" db:"product_id"`
	Quantity  int64 `json:"quantity" binding:"required" db:"quantity"`
}
