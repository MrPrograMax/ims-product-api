package product_ser

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"ims-product-api/model"
	"ims-product-api/pkg/repository"
)

type ProductService struct {
	repo         repository.Product
	categoryRepo repository.Category
	locationRepo repository.Location
	statusRepo   repository.ProductStatus
}

func NewProductService(pr repository.Product, cat repository.Category, loc repository.Location, ps repository.ProductStatus) *ProductService {
	return &ProductService{
		repo:         pr,
		categoryRepo: cat,
		locationRepo: loc,
		statusRepo:   ps,
	}
}

func (s *ProductService) Create(product model.Product) (int64, error) {
	return s.repo.Create(product)
}

func (s *ProductService) Update(id int64, product model.UpdateProduct) error {
	if err := model.Verify(product); err != nil {
		logrus.Error(err)
		return err
	}

	pr, err := s.repo.GetById(id)
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("product with specified id not found")
	}

	if dtoName := product.Name; dtoName != nil && *dtoName != pr.Name {
		if _, err := s.repo.GetByName(*dtoName); err != nil {
			logrus.Error(err)
			return fmt.Errorf("product with such name already exists")
		}
	}

	return s.repo.Update(pr.Id, product)
}

func (s *ProductService) Delete(id int64) error {
	if _, err := s.repo.GetById(id); err != nil {
		logrus.Error(err)
		return fmt.Errorf("product with specified id not found")
	}

	return s.repo.Delete(id)
}

func (s *ProductService) GetAll() ([]model.ProductDTO, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("products not found")
	}

	return mapProductToDtoList(s, products)
}

func (s *ProductService) GetById(id int64) (model.ProductDTO, error) {
	product, err := s.repo.GetById(id)
	if err != nil {
		logrus.Error(err)
		return model.ProductDTO{}, fmt.Errorf("product not found")
	}

	return mapProductToDto(s, product)
}

func (s *ProductService) GetByName(name string) (model.ProductDTO, error) {
	product, err := s.repo.GetByName(name)
	if err != nil {
		logrus.Error(err)
		return model.ProductDTO{}, fmt.Errorf("product not found")
	}

	return mapProductToDto(s, product)
}

func (s *ProductService) GetByCategoryName(name string) ([]model.ProductDTO, error) {
	cat, err := s.categoryRepo.GetByName(name)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("no such category found")
	}

	products, err := s.repo.GetByCategoryId(cat.Id)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("no such products found")
	}

	return mapProductToDtoList(s, products)
}

func (s *ProductService) GetByLocation(location model.Location) ([]model.ProductDTO, error) {
	loc, err := s.locationRepo.GetByRowAndPlace(location.Row, location.Place)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("no such location found")
	}

	products, err := s.repo.GetByCategoryId(loc.Id)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("no such products found")
	}

	return mapProductToDtoList(s, products)
}

func (s *ProductService) GetByStatusName(name string) ([]model.ProductDTO, error) {
	stat, err := s.statusRepo.GetByName(name)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("no such status found")
	}

	products, err := s.repo.GetByStatusId(stat.Id)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("no such products found")
	}

	return mapProductToDtoList(s, products)
}

func (s *ProductService) GetByCategoryId(id int64) ([]model.ProductDTO, error) {
	products, err := s.repo.GetByCategoryId(id)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("no such products found")
	}

	return mapProductToDtoList(s, products)
}

func (s *ProductService) GetByLocationId(id int64) ([]model.ProductDTO, error) {
	products, err := s.repo.GetByLocationId(id)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("no such products found")
	}

	return mapProductToDtoList(s, products)
}

func (s *ProductService) GetByStatusId(id int64) ([]model.ProductDTO, error) {
	products, err := s.repo.GetByStatusId(id)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("no such products found")
	}

	return mapProductToDtoList(s, products)
}

func mapProductToDtoList(s *ProductService, products []model.Product) ([]model.ProductDTO, error) {
	categories, err := s.categoryRepo.GetAll()
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("error getting categories in products")
	}

	locations, err := s.locationRepo.GetAll()
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("error getting locations in products")
	}

	statuses, err := s.statusRepo.GetAll()
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("error getting statuses in products")
	}

	dtos := make([]model.ProductDTO, len(products))
	for i, product := range products {
		cat, err := findCategoryById(categories, products[i].CategoryId)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		loc, err := findLocationById(locations, products[i].LocationId)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		stat, err := findStatusById(statuses, products[i].StatusId)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		dtos[i] = model.ProductDTO{
			Id:          product.Id,
			Name:        product.Name,
			Quantity:    product.Quantity,
			Description: product.Description,
			Category:    cat,
			Location:    loc,
			Status:      stat,
		}
	}
	return dtos, nil
}

func mapProductToDto(s *ProductService, product model.Product) (model.ProductDTO, error) {
	cat, err := s.categoryRepo.GetById(product.CategoryId)
	if err != nil {
		logrus.Error(err)
		return model.ProductDTO{}, fmt.Errorf("no such category found")
	}

	loc, err := s.locationRepo.GetById(product.LocationId)
	if err != nil {
		logrus.Error(err)
		return model.ProductDTO{}, fmt.Errorf("no such location found")
	}

	stat, err := s.statusRepo.GetById(product.StatusId)
	if err != nil {
		logrus.Error(err)
		return model.ProductDTO{}, fmt.Errorf("no such status found")
	}

	return model.ProductDTO{
		Id:          product.Id,
		Name:        product.Name,
		Quantity:    product.Quantity,
		Description: product.Description,
		Category:    cat,
		Location:    loc,
		Status:      stat,
	}, nil
}

func findCategoryById(cats []model.Category, id int64) (model.Category, error) {
	for _, cat := range cats {
		logrus.Info(cat.Id, " - ", id)
		if cat.Id == id {
			return cat, nil
		}
	}
	return model.Category{}, fmt.Errorf("category not found")
}

func findLocationById(locs []model.Location, id int64) (model.Location, error) {
	for _, loc := range locs {
		if loc.Id == id {
			return loc, nil
		}
	}
	return model.Location{}, fmt.Errorf("location not found")
}

func findStatusById(statuses []model.ProductStatus, id int64) (model.ProductStatus, error) {
	for _, status := range statuses {
		if status.Id == id {
			return status, nil
		}
	}
	return model.ProductStatus{}, fmt.Errorf("status not found")
}
