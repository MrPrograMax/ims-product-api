package order_rep

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"ims-product-api/model"
	"strings"
)

const orderItemTable = "order_item"

var orderItemFields = []string {
	"order_id",
	"product_id",
	"quantity_id",
}

type OrderItem struct {
	db *sqlx.DB
}

func NewOrderItem(db *sqlx.DB) *OrderItem {
	return &OrderItem{
		db: db,
	}
}

func (r *OrderItem) Create(item model.OrderItem) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	id, err2 := createItem(item, tx)
	if err2 != nil {
		logrus.Error(err)
		if err := tx.Rollback(); err != nil {
			logrus.Error(err)
			return 0, err
		}
		return 0, err2
	}

	return id, tx.Commit()
}

func (r *OrderItem) CreateList(items []model.OrderItem) ([]int64, error) {
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

func (r *OrderItem) Delete(id int64) error {
	deleteQuery := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, orderItemTable)

	_, err := r.db.Exec(deleteQuery, id)
	return err
}

func (r *OrderItem) GetAll() ([]model.OrderItem, error) {
	var items []model.OrderItem
	query := fmt.Sprintf(`SELECT * FROM %s`, orderItemTable)

	if err := r.db.Select(&items, query); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return items, nil
}

func (r *OrderItem) GetById(id int64) (model.OrderItem, error) {
	var item model.OrderItem
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id=$1`, orderItemTable)

	if err := r.db.Select(&item, query, id); err != nil {
		logrus.Error(err)
		return item, err
	}

	return item, nil
}

func (r *OrderItem) GetByOrderId(id int64) ([]model.OrderItem, error) {
	var items []model.OrderItem
	query := fmt.Sprintf(`SELECT * FROM %s WHERE order_id=$1`, orderItemTable)

	if err := r.db.Select(&items, query, id); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return items, nil
}

func (r *OrderItem) GetByProductId(id int64) ([]model.OrderItem, error) {
	var items []model.OrderItem
	query := fmt.Sprintf(`SELECT * FROM %s WHERE product_id=$1`, orderItemTable)

	if err := r.db.Select(&items, query, id); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return items, nil
}

func createItem(item model.OrderItem, tx *sql.Tx) (int64, error) {
	var id int64
	insertQuery := fmt.Sprintf(`INSERT INTO %s (%s) VALUES ($1, $2, $3) RETURNING id`,
		orderItemTable, strings.Join(orderItemFields, ", "))

	row := tx.QueryRow(insertQuery, item.OrderId, item.ProductId, item.Quantity)

	if err := row.Scan(&id); err != nil {
		logrus.Error(err)
		return 0, err
	}
	return id, nil
}
