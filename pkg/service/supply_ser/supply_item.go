package supply_ser

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"ims-product-api/model"
	"ims-product-api/pkg/repository"
)

type SupplyItemService struct {
	repo repository.SupplyItem
	prodRepo repository.Product
}

func NewSupplyItemService(repo repository.SupplyItem, prod repository.Product) *SupplyItemService {
	return &SupplyItemService{
		repo: repo,
		prodRepo: prod,
	}
}

func (s *SupplyItemService) Create(item model.SupplyItem) (int64, error) {
	prod, err := s.prodRepo.GetById(item.ProductId)
	if err != nil {
		logrus.Error(err)
		return 0, fmt.Errorf("product with supply item id not found")
	}

	newQuantity := prod.Quantity + item.Quantity
	if err := updateProduct(s, item.ProductId, newQuantity); err != nil {
		logrus.Error(err)
		return 0, err
	}

	return s.repo.Create(item)
}

func (s *SupplyItemService) CreateList(items []model.SupplyItem) ([]int64, error) {
	var newQuantities map[int64]int64
	for _, item := range items {
		prod, err := s.prodRepo.GetById(item.ProductId)
		if err != nil {
			logrus.Error(err)
			return []int64{}, fmt.Errorf("product in supply item id not found")
		}

		newQuantities[item.ProductId] = prod.Quantity + item.Quantity
	}

	for k, v := range newQuantities {
		if err := updateProduct(s, k, v); err != nil {
			return []int64{}, err
		}
	}

	return s.repo.CreateList(items)
}

func (s *SupplyItemService) Delete(id int64) error {
	item, err := s.repo.GetById(id)
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("supply item with specified id not found")
	}

	prod, err := s.prodRepo.GetById(item.ProductId)
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("product in supply item id not found")
	}

	newQuantity := prod.Quantity - item.Quantity
	if newQuantity < 0 {
		newQuantity = 0
	}

	if err := s.prodRepo.Update(item.ProductId, model.UpdateProduct{
		Quantity: &(newQuantity),
	}); err != nil {
		logrus.Error(err)
		return fmt.Errorf("unable to make update of product")
	}

	return s.repo.Delete(id)
}

func (s *SupplyItemService) GetAll() ([]model.SupplyItem, error) {
	return s.repo.GetAll()
}

func (s *SupplyItemService) GetById(id int64) (model.SupplyItem, error) {
	return s.repo.GetById(id)
}

func (s *SupplyItemService) GetBySupplyId(id int64) ([]model.SupplyItem, error) {
	return s.repo.GetBySupplyId(id)
}

func (s *SupplyItemService) GetByProductId(id int64) ([]model.SupplyItem, error) {
	return s.repo.GetByProductId(id)
}

func updateProduct(s *SupplyItemService, prodId int64, newQuantity int64) error {
	if err := s.prodRepo.Update(prodId, model.UpdateProduct{
		Quantity: &(newQuantity),
	}); err != nil {
		return fmt.Errorf("unable to make update of product")
	}
	return nil
}
