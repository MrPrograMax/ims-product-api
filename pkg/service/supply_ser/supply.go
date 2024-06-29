package supply_ser

import (
	"ims-product-api/model"
	"ims-product-api/pkg/repository"
)

type SupplyService struct {
	repo repository.Supply
}

func NewSupplyService(repo repository.Supply) *SupplyService {
	return &SupplyService{
		repo: repo,
	}
}

func (s *SupplyService) Create() (int64, error) {
	return s.repo.Create()
}

func (s *SupplyService) Delete(id int64) error {
	return s.repo.Delete(id)
}

func (s *SupplyService) GetAll() ([]model.Supply, error) {
	return s.repo.GetAll()
}
