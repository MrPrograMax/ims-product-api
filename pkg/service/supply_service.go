package service

import (
	"ims-product-api/model"
	"ims-product-api/pkg/repository"
	"ims-product-api/pkg/service/supply_ser"
)

type Supply interface {
	Create() (int64, error)
	Delete(id int64) error

	GetAll() ([]model.Supply, error)
}

type SupplyItem interface {
	Create(item model.SupplyItem) (int64, error)
	CreateList(items []model.SupplyItem) ([]int64, error)
	Delete(id int64) error

	GetAll() ([]model.SupplyItem, error)
	GetById(id int64) (model.SupplyItem, error)
	GetBySupplyId(id int64) ([]model.SupplyItem, error)
	GetByProductId(id int64) ([]model.SupplyItem, error)
}

type SupplyService struct {
	Supply
	SupplyItem
}

func NewSupplyService(repos *repository.SupplyRepository, prod *repository.ProductRepository) *SupplyService {
	return &SupplyService{
		Supply: supply_ser.NewSupplyService(repos.Supply),
		SupplyItem: supply_ser.NewSupplyItemService(repos.SupplyItem, prod.Product),
	}
}
