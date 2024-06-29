package product_rep

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"ims-product-api/model"
	"reflect"
	"strings"
)

const productTable = "product"

var productFields = []string{
	`"name"`,
	"quantity",
	"description",
	"category_id",
	"location_id",
	"status_id",
}

type Product struct {
	db *sqlx.DB
}

func NewProduct(db *sqlx.DB) *Product {
	return &Product{
		db: db,
	}
}

func (p *Product) Create(product model.Product) (int64, error) {
	tx, err := p.db.Begin()
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	var id int64
	insertQuery := fmt.Sprintf(`INSERT INTO %s (%s) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		productTable, strings.Join(productFields, ", "))

	row := tx.QueryRow(insertQuery, product.Name, product.Quantity, product.Description,
		product.CategoryId, product.LocationId, product.StatusId)

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

func (p *Product) Update(id int64, product model.UpdateProduct) error {
	placeholders := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	prodType := reflect.ValueOf(product).Elem()
	for i := 0; i < prodType.NumField(); i++ {
		valueField := prodType.Field(i)
		typeField := prodType.Type().Field(i)
		if valueField.Kind() == reflect.Ptr && !valueField.IsNil() {
			dbName := typeField.Tag.Get("db")
			if dbName == "" {
				dbName = typeField.Name
			}

			placeholders = append(placeholders, fmt.Sprintf("%s=$%d", dbName, argId))
			args = append(args, valueField.Elem())
			argId++
		}
	}

	//another realization

	//if name := product_rep.Name; name != nil {
	//	placeholders = append(placeholders, fmt.Sprintf("name=$%d", argId))
	//	args = append(args, *name)
	//	argId++
	//}
	//
	//if description := product_rep.Description; description != nil {
	//	placeholders = append(placeholders, fmt.Sprintf("description=$%d", argId))
	//	args = append(args, *description)
	//	argId++
	//}
	//
	//if quantity := product_rep.Quantity; quantity != nil {
	//	placeholders = append(placeholders, fmt.Sprintf("quantity=$%d", argId))
	//	args = append(args, *quantity)
	//	argId++
	//}
	//
	//if category := product_rep.CategoryId; category != nil {
	//	placeholders = append(placeholders, fmt.Sprintf("category_id=$%d", argId))
	//	args = append(args, *category)
	//	argId++
	//}
	//
	//if location := product_rep.LocationId; location != nil {
	//	placeholders = append(placeholders, fmt.Sprintf("location_id=$%d", argId))
	//	args = append(args, *location)
	//	argId++
	//}
	//
	//if status := product_rep.StatusId; status != nil {
	//	placeholders = append(placeholders, fmt.Sprintf("status_id=$%d", argId))
	//	args = append(args, *status)
	//	argId++
	//}

	if argId == 1 {
		return nil
	}

	updateString := strings.Join(placeholders, ", ")
	updateQuery := fmt.Sprintf(`UPDATE %s SET %s WHERE id=$%d`, productTable, updateString, argId)
	args = append(args, id)

	_, err := p.db.Exec(updateQuery, args...)
	return err
}

func (p *Product) Delete(id int64) error {
	deleteQuery := fmt.Sprintf(`UPDATE %s SET status_id=(SELECT id
          														FROM %s
          														WHERE "name"='inactive')
          								WHERE id=$1`,
		productTable, productsStatusTable)

	_, err := p.db.Exec(deleteQuery, id)
	return err
}

func (p *Product) GetAll() ([]model.Product, error) {
	var products []model.Product
	query := fmt.Sprintf(`SELECT * FROM %s`, productTable)

	if err := p.db.Select(&products, query); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return products, nil
}

func (p *Product) GetById(id int64) (model.Product, error) {
	var product model.Product
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id=$1`, productTable)

	if err := p.db.Select(&product, query, id); err != nil {
		logrus.Error(err)
		return product, err
	}

	return product, nil
}

func (p *Product) GetByName(name string) (model.Product, error) {
	var product model.Product
	query := fmt.Sprintf(`SELECT * FROM %s WHERE "name"=$1`, productTable)

	if err := p.db.Select(&product, query, name); err != nil {
		logrus.Error(err)
		return product, err
	}

	return product, nil
}

func (p *Product) GetByCategoryId(id int64) ([]model.Product, error) {
	var products []model.Product
	query := fmt.Sprintf(`SELECT * FROM %s WHERE category_id=$1`, productTable)

	if err := p.db.Select(&products, query, id); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return products, nil
}

func (p *Product) GetByLocationId(id int64) ([]model.Product, error) {
	var products []model.Product
	query := fmt.Sprintf(`SELECT * FROM %s WHERE location_id=$1`, productTable)

	if err := p.db.Select(&products, query, id); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return products, nil
}

func (p *Product) GetByStatusId(id int64) ([]model.Product, error) {
	var products []model.Product
	query := fmt.Sprintf(`SELECT * FROM %s WHERE status_id=$1`, productTable)

	if err := p.db.Select(&products, query, id); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return products, nil
}
