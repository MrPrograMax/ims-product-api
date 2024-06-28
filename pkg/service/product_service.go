package service

import (
	"ims-product-api/model"
	"ims-product-api/pkg/repository"
	"ims-product-api/pkg/service/product_ser"
)

type Product interface {
	Create(product model.Product) (int64, error)
	Update(id int64, product model.UpdateProduct) error
	Delete(id int64) error

	GetAll() ([]model.ProductDTO, error)

	GetByName(name string) (model.ProductDTO, error)
	GetByCategoryName(name string) ([]model.ProductDTO, error)
	GetByLocation(location model.Location) ([]model.ProductDTO, error)
	GetByStatusName(name string) ([]model.ProductDTO, error)

	GetById(id int64) (model.ProductDTO, error)
	GetByCategoryId(id int64) ([]model.ProductDTO, error)
	GetByLocationId(id int64) ([]model.ProductDTO, error)
	GetByStatusId(id int64) ([]model.ProductDTO, error)
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
	GetByRowAndPlace(row string, place string) (model.Location, error)
	GetByRow(row string) ([]model.Location, error)
}

type ProductStatus interface {
	GetAll() ([]model.ProductStatus, error)
	GetByName(name string) (model.ProductStatus, error)
	GetById(id int64) (model.ProductStatus, error)
}

type ProductService struct {
	Product
	Category
	Location
	ProductStatus
}

func NewProductService(repos *repository.ProductRepository) *ProductService {
	return &ProductService{
		Product:       product_ser.NewProductService(repos.Product, repos.Category, repos.Location, repos.ProductStatus),
		Category:      product_ser.NewCategoryService(repos.Category),
		Location:      product_ser.NewLocationService(repos.Location),
		ProductStatus: product_ser.NewProductStatusService(repos.ProductStatus),
	}
}
