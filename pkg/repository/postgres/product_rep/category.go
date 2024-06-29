package product_rep

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"ims-product-api/model"
	"strings"
)

const categoryTable = "category"

var categoryFields = []string{
	`"name"`,
}

type Category struct {
	db *sqlx.DB
}

func NewCategory(db *sqlx.DB) *Category {
	return &Category{
		db: db,
	}
}

func (cat *Category) Create(category model.Category) (int64, error) {
	tx, err := cat.db.Begin()
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	var id int64
	insertQuery := fmt.Sprintf(`INSERT INTO %s (%s) VALUES ($1) RETURNING id`,
		categoryTable, strings.Join(categoryFields, ", "))

	row := tx.QueryRow(insertQuery, category.Name)

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

func (cat *Category) Update(id int64, category model.UpdateCategory) error {
	if name := category.Name; name != nil {
		updateQuery := fmt.Sprintf(`UPDATE %s SET "name"=$1 WHERE id=$2`, categoryTable)

		_, err := cat.db.Exec(updateQuery, *name, id)
		return err
	}

	return nil
}

func (cat *Category) Delete(id int64) error {
	deleteQuery := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, categoryTable)

	_, err := cat.db.Exec(deleteQuery, id)
	return err
}

func (cat *Category) GetAll() ([]model.Category, error) {
	var categories []model.Category
	query := fmt.Sprintf(`SELECT * FROM %s`, categoryTable)

	if err := cat.db.Select(&categories, query); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return categories, nil
}

func (cat *Category) GetById(id int64) (model.Category, error) {
	var category model.Category
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id=$1`, categoryTable)

	if err := cat.db.Select(&category, query, id); err != nil {
		logrus.Error(err)
		return category, err
	}

	return category, nil
}

func (cat *Category) GetByName(name string) (model.Category, error) {
	var category model.Category
	query := fmt.Sprintf(`SELECT * FROM %s WHERE "name"=$1`, categoryTable)

	if err := cat.db.Select(&category, query, name); err != nil {
		logrus.Error(err)
		return category, err
	}

	return category, nil
}
