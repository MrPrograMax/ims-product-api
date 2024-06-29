package product_rep

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"ims-product-api/model"
	"strings"
)

const locationTable = "location"

var locationFields = []string{
	`"row"`,
	"place",
}

type Location struct {
	db *sqlx.DB
}

func NewLocation(db *sqlx.DB) *Location {
	return &Location{
		db: db,
	}
}

func (loc *Location) Create(location model.Location) (int64, error) {
	tx, err := loc.db.Begin()
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	var id int64
	insertQuery := fmt.Sprintf(`INSERT INTO %s (%s) VALUES ($1, $2) RETURNING id`,
		locationTable, strings.Join(locationFields, ", "))

	row := tx.QueryRow(insertQuery, location.Row, location.Place)

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

func (loc *Location) Update(id int64, location model.UpdateLocation) error {
	placeholders := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if row := location.Row; row != nil {
		placeholders = append(placeholders, fmt.Sprintf(`"row"=$%d`, argId))
		args = append(args, *row)
		argId++
	}

	if place := location.Place; place != nil {
		placeholders = append(placeholders, fmt.Sprintf("place=$%d", argId))
		args = append(args, *place)
		argId++
	}

	if argId == 1 {
		return nil
	}

	updateString := strings.Join(placeholders, ", ")
	updateQuery := fmt.Sprintf(`UPDATE %s SET %s WHERE id=$%d`, locationTable, updateString, argId)
	args = append(args, id)

	_, err := loc.db.Exec(updateQuery, args...)
	return err
}

func (loc *Location) Delete(id int64) error {
	deleteQuery := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, locationTable)

	_, err := loc.db.Exec(deleteQuery, id)
	return err
}

func (loc *Location) GetAll() ([]model.Location, error) {
	var locations []model.Location
	query := fmt.Sprintf(`SELECT * FROM %s`, locationTable)

	if err := loc.db.Select(&locations, query); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return locations, nil
}

func (loc *Location) GetById(id int64) (model.Location, error) {
	var location model.Location
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id=$1`, locationTable)

	if err := loc.db.Select(&location, query, id); err != nil {
		logrus.Error(err)
		return location, err
	}

	return location, nil
}

func (loc *Location) GetListByRow(row string) ([]model.Location, error) {
	var locations []model.Location
	query := fmt.Sprintf(`SELECT * FROM %s WHERE "row"=$1`, locationTable)

	if err := loc.db.Select(&locations, query, row); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return locations, nil
}

func (loc *Location) GetByRowAndPlace(row string, place string) (model.Location, error) {
	var location model.Location
	query := fmt.Sprintf(`SELECT * FROM %s WHERE "row"=$1 AND place=$2`, locationTable)

	if err := loc.db.Select(&location, query, row, place); err != nil {
		logrus.Error(err)
		return location, err
	}

	return location, nil
}
