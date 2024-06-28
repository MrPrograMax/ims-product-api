package product_rep

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"ims-product-api/model"
)

const productsStatusTable = "products_status"

type ProductStatus struct {
	db *sqlx.DB
}

func NewProductStatus(db *sqlx.DB) *ProductStatus {
	return &ProductStatus{
		db: db,
	}
}

func (ps *ProductStatus) GetAll() ([]model.ProductStatus, error) {
	var status []model.ProductStatus
	query := fmt.Sprintf(`SELECT * FROM %s`, productsStatusTable)

	if err := ps.db.Select(&status, query); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return status, nil
}

func (ps *ProductStatus) GetByName(name string) (model.ProductStatus, error) {
	var status model.ProductStatus
	query := fmt.Sprintf(`SELECT * FROM %s WHERE "name"=$1`, productsStatusTable)

	if err := ps.db.Select(&status, query, name); err != nil {
		logrus.Error(err)
		return status, err
	}

	return status, nil
}

func (ps *ProductStatus) GetById(id int64) (model.ProductStatus, error) {
	var status model.ProductStatus
	query := fmt.Sprintf(`SELECT * FROM %s WHERE "id"=$1`, productsStatusTable)

	if err := ps.db.Select(&status, query, id); err != nil {
		logrus.Error(err)
		return status, err
	}

	return status, nil
}
