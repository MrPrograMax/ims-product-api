package repository

import (
	"github.com/jmoiron/sqlx"
	"ims-product-api/model"
	"ims-product-api/pkg/repository/product_rep"
)

type Product interface {
	Create(product model.Product) (int64, error)
	Update(id int64, product model.UpdateProduct) error
	Delete(id int64) error

	GetAll() ([]model.Product, error)
	GetById(id int64) (model.Product, error)
	GetByName(name string) (model.Product, error)
	GetByCategoryId(id int64) ([]model.Product, error)
	GetByLocationId(id int64) ([]model.Product, error)
	GetByStatusId(id int64) ([]model.Product, error)
}

type Category interface {
	Create(category model.Category) (int64, error)
	Update(id int64, category model.UpdateCategory) error
	Delete(id int64) error

	GetAll() ([]model.Category, error)
	GetById(id int64) (model.Category, error)
	GetByName(name string) (model.Category, error)
}

type Location interface {
	Create(location model.Location) (int64, error)
	Update(id int64, location model.UpdateLocation) error
	Delete(id int64) error

	GetAll() ([]model.Location, error)
	GetById(id int64) (model.Location, error)
	GetListByRow(row string) ([]model.Location, error)
	GetByRowAndPlace(row string, place string) (model.Location, error)
}

type ProductStatus interface {
	GetAll() ([]model.ProductStatus, error)
	GetByName(name string) (model.ProductStatus, error)
	GetById(id int64) (model.ProductStatus, error)
}

type ProductRepository struct {
	Product
	Category
	Location
	ProductStatus
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		Product:       product_rep.NewProduct(db),
		Category:      product_rep.NewCategory(db),
		Location:      product_rep.NewLocation(db),
		ProductStatus: product_rep.NewProductStatus(db),
	}
}
