package order_ser

import (
	"ims-product-api/model"
	"ims-product-api/pkg/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService) Create() (int64, error) {
	return s.repo.Create()
}

func (s *OrderService) Delete(id int64) error {
	return s.repo.Delete(id)
}

func (s *OrderService) GetAll() ([]model.Order, error) {
	return s.repo.GetAll()
}
