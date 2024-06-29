package supply_rep

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"ims-product-api/model"
	"time"
)

const supplyTable = `"order"`

const supplyField = "datetime"

type Supply struct {
	db *sqlx.DB
}

func NewSupply(db *sqlx.DB) *Supply {
	return &Supply{
		db: db,
	}
}

func (r *Supply) Create() (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	var id int64
	insertQuery := fmt.Sprintf(`INSERT INTO %s (%s) VALUES ($1) RETURNING id`,
		supplyTable, supplyField)

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

func (r *Supply) Delete(id int64) error {
	deleteQuery := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, supplyTable)

	_, err := r.db.Exec(deleteQuery, id)
	return err
}

func (r *Supply) GetAll() ([]model.Supply, error) {
	var supplies []model.Supply
	query := fmt.Sprintf(`SELECT * FROM %s`, supplyTable)

	if err := r.db.Select(&supplies, query); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return supplies, nil
}
