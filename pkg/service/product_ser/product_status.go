package product_ser

import (
	"ims-product-api/model"
	"ims-product-api/pkg/repository"
)

type ProductStatusService struct {
	repo repository.ProductStatus
}

func NewProductStatusService(repo repository.ProductStatus) *ProductStatusService {
	return &ProductStatusService{
		repo: repo,
	}
}

func (s *ProductStatusService) GetAll() ([]model.ProductStatus, error) {
	return s.repo.GetAll()
}

func (s *ProductStatusService) GetByName(name string) (model.ProductStatus, error) {
	return s.repo.GetByName(name)
}

func (s *ProductStatusService) GetById(id int64) (model.ProductStatus, error) {
	return s.repo.GetById(id)
}
