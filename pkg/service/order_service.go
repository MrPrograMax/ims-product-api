package service

import (
	"ims-product-api/model"
	"ims-product-api/pkg/repository"
	"ims-product-api/pkg/service/order_ser"
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

type OrderService struct {
	Order
	OrderItem
}

func NewOrderService(repos *repository.OrderRepository, prod *repository.ProductRepository) *OrderService {
	return &OrderService{
		Order: order_ser.NewOrderService(repos.Order),
		OrderItem: order_ser.NewOrderItemService(repos.OrderItem, prod.Product),
	}
}
