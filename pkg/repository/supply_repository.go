package repository

import (
	"github.com/jmoiron/sqlx"
	"ims-product-api/model"
	"ims-product-api/pkg/repository/postgres/supply_rep"
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

type SupplyRepository struct {
	Supply
	SupplyItem
}

func NewSupplyRepository(db *sqlx.DB) *SupplyRepository {
	return &SupplyRepository{
		Supply: supply_rep.NewSupply(db),
		SupplyItem: supply_rep.NewSupplyItem(db),
	}
}
