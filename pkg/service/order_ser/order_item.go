package order_ser

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"ims-product-api/model"
	"ims-product-api/pkg/repository"
)

type OrderItemService struct {
	repo     repository.OrderItem
	prodRepo repository.Product
}

func NewOrderItemService(repo repository.OrderItem, prod repository.Product) *OrderItemService {
	return &OrderItemService{
		repo:     repo,
		prodRepo: prod,
	}
}

func (s *OrderItemService) Create(item model.OrderItem) (int64, error) {
	prod, err := s.prodRepo.GetById(item.ProductId)
	if err != nil {
		logrus.Error(err)
		return 0, fmt.Errorf("product in order item id not found")
	}

	if prod.Quantity < item.Quantity {
		logrus.Error(err)
		return 0, fmt.Errorf("quantity of product less than in order")
	}

	newQuantity := prod.Quantity - item.Quantity
	if err := updateProduct(s, item.ProductId, newQuantity); err != nil {
		return 0, err
	}

	return s.repo.Create(item)
}

func (s *OrderItemService) CreateList(items []model.OrderItem) ([]int64, error) {
	var newQuantities map[int64]int64
	for _, item := range items {
		prod, err := s.prodRepo.GetById(item.ProductId)
		if err != nil {
			logrus.Error(err)
			return []int64{}, fmt.Errorf("product in order item id not found")
		}

		if prod.Quantity < item.Quantity {
			logrus.Error(err)
			return []int64{}, fmt.Errorf("quantity of product less than in order")
		}

		newQuantities[item.ProductId] = prod.Quantity - item.Quantity
	}

	for k, v := range newQuantities {
		if err := updateProduct(s, k, v); err != nil {
			return []int64{}, err
		}
	}

	return s.repo.CreateList(items)
}

func (s *OrderItemService) Delete(id int64) error {
	item, err := s.repo.GetById(id)
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("order item with specified id not found")
	}

	prod, err := s.prodRepo.GetById(item.ProductId)
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("product in order item id not found")
	}

	newQuantity := prod.Quantity + item.Quantity
	if err := s.prodRepo.Update(item.ProductId, model.UpdateProduct{
		Quantity: &(newQuantity),
	}); err != nil {
		logrus.Error(err)
		return fmt.Errorf("unable to make update of product")
	}

	return s.repo.Delete(id)
}

func (s *OrderItemService) GetAll() ([]model.OrderItem, error) {
	return s.repo.GetAll()
}

func (s *OrderItemService) GetById(id int64) (model.OrderItem, error) {
	return s.repo.GetById(id)
}

func (s *OrderItemService) GetByOrderId(id int64) ([]model.OrderItem, error) {
	return s.repo.GetByOrderId(id)
}

func (s *OrderItemService) GetByProductId(id int64) ([]model.OrderItem, error) {
	return s.repo.GetByProductId(id)
}

func updateProduct(s *OrderItemService, prodId int64, newQuantity int64) error {
	if err := s.prodRepo.Update(prodId, model.UpdateProduct{
		Quantity: &(newQuantity),
	}); err != nil {
		return fmt.Errorf("unable to make update of product")
	}
	return nil
}
