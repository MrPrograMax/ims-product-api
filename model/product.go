package model

type Product struct {
	Id          int64  `json:"id" db:"id"`
	Name        string `json:"name" binding:"required" db:"name"`
	Quantity    int64  `json:"quantity" binding:"required" db:"quantity"`
	Description string `json:"description" binding:"required" db:"description"`
	CategoryId  int64  `json:"category" binding:"required" db:"category_id"`
	LocationId  int64  `json:"location" binding:"required" db:"location_id"`
	StatusId    int64  `json:"status" binding:"required" db:"status_id"`
}

type Category struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"name" binding:"required" db:"name"`
}

type Location struct {
	Id    int64  `json:"id" db:"id"`
	Row   string `json:"row" binding:"required" db:"row"`
	Place string `json:"place" binding:"required" db:"place"`
}

type ProductStatus struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"name" binding:"required" db:"name"`
}
