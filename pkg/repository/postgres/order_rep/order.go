package order_rep

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"ims-product-api/model"
	"time"
)

const orderTable = `"order"`

const orderField = "datetime"

type Order struct {
	db *sqlx.DB
}

func NewOrder(db *sqlx.DB) *Order {
	return &Order{
		db: db,
	}
}

func (r *Order) Create() (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	var id int64
	insertQuery := fmt.Sprintf(`INSERT INTO %s (%s) VALUES ($1) RETURNING id`,
		orderTable, orderField)

	row := tx.QueryRow(insertQuery, time.Now())

	if err = row.Scan(&id); err != nil {
		logrus.Error(err)
		if err := tx.Rollback(); err != nil {
			logrus.Error(err)
			return 0, err
		}
		return 0, err
	}

	return id, tx.Commit()
}

func (r *Order) Delete(id int64) error {
	deleteQuery := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, orderTable)

	_, err := r.db.Exec(deleteQuery, id)
	return err
}

func (r *Order) GetAll() ([]model.Order, error) {
	var orders []model.Order
	query := fmt.Sprintf(`SELECT * FROM %s`, orderTable)

	if err := r.db.Select(&orders, query); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return orders, nil
}
