package supply_rep

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"ims-product-api/model"
	"strings"
)

const supplyItemTable = "order_item"

var supplyItemFields = []string {
	"order_id",
	"product_id",
	"quantity_id",
}

type SupplyItem struct {
	db *sqlx.DB
}

func NewSupplyItem(db *sqlx.DB) *SupplyItem {
	return &SupplyItem{
		db: db,
	}
}

func (r *SupplyItem) Create(item model.SupplyItem) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	id, err := createItem(item, tx)
	if err != nil {
		logrus.Error(err)
		if err := tx.Rollback(); err != nil {
			logrus.Error(err)
			return 0, err
		}
		return 0, err
	}

	return id, tx.Commit()
}

func (r *SupplyItem) CreateList(items []model.SupplyItem) ([]int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		logrus.Error(err)
		return []int64{}, err
	}

	ids := make([]int64, len(items))
	for i, item := range items {
		id, err := createItem(item, tx)
		if err != nil {
			logrus.Error(err)
			if err := tx.Rollback(); err != nil {
				logrus.Error(err)
				return []int64{}, err
			}

			return []int64{}, err
		}

		ids[i] = id
	}

	return ids, tx.Commit()
}

func (r *SupplyItem) Delete(id int64) error {
	deleteQuery := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, supplyItemTable)

	_, err := r.db.Exec(deleteQuery, id)
	return err
}

func (r *SupplyItem) GetAll() ([]model.SupplyItem, error) {
	var items []model.SupplyItem
	query := fmt.Sprintf(`SELECT * FROM %s`, supplyItemTable)

	if err := r.db.Select(&items, query); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return items, nil
}

func (r *SupplyItem) GetById(id int64) (model.SupplyItem, error) {
	var item model.SupplyItem
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id=$1`, supplyItemTable)

	if err := r.db.Select(&item, query, id); err != nil {
		logrus.Error(err)
		return item, err
	}

	return item, nil
}

func (r *SupplyItem) GetBySupplyId(id int64) ([]model.SupplyItem, error) {
	var items []model.SupplyItem
	query := fmt.Sprintf(`SELECT * FROM %s WHERE supply_id=$1`, supplyItemTable)

	if err := r.db.Select(&items, query, id); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return items, nil
}

func (r *SupplyItem) GetByProductId(id int64) ([]model.SupplyItem, error) {
	var items []model.SupplyItem
	query := fmt.Sprintf(`SELECT * FROM %s WHERE product_id=$1`, supplyItemTable)

	if err := r.db.Select(&items, query, id); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return items, nil
}

func createItem(item model.SupplyItem, tx *sql.Tx) (int64, error) {
	var id int64
	insertQuery := fmt.Sprintf(`INSERT INTO %s (%s) VALUES ($1, $2, $3) RETURNING id`,
		supplyItemTable, strings.Join(supplyItemFields, ", "))

	row := tx.QueryRow(insertQuery, item.SupplyId, item.ProductId, item.Quantity)

	if err := row.Scan(&id); err != nil {
		logrus.Error(err)
		return 0, err
	}
	return id, nil
}
