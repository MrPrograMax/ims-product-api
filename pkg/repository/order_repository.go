package repository

import (
	"github.com/jmoiron/sqlx"
	"ims-product-api/model"
	"ims-product-api/pkg/repository/postgres/order_rep"
)

type Order interface {
	Create() (int64, error)
	Delete(id int64) error

	GetAll() ([]model.Order, error)
}

type OrderItem interface {
	Create(item model.OrderItem) (int64, error)
	CreateList(items []model.OrderItem) ([]int64, error)
	Delete(id int64) error

	GetAll() ([]model.OrderItem, error)
	GetById(id int64) (model.OrderItem, error)
	GetByOrderId(id int64) ([]model.OrderItem, error)
	GetByProductId(id int64) ([]model.OrderItem, error)
}

type OrderRepository struct {
	Order
	OrderItem
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{
		Order: order_rep.NewOrder(db),
		OrderItem: order_rep.NewOrderItem(db),
	}
}
