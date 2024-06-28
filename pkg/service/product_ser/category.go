package product_ser

import (
	"github.com/sirupsen/logrus"
	"ims-product-api/model"
	"ims-product-api/pkg/repository"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) Create(category model.Category) (int64, error) {
	return s.repo.Create(category)
}

func (s *CategoryService) Update(id int64, category model.UpdateCategory) error {
	if err := model.Verify(category); err != nil {
		logrus.Error(err)
		return err
	}

	return s.repo.Update(id, category)
}

func (s *CategoryService) Delete(id int64) error {
	return s.repo.Delete(id)
}

func (s *CategoryService) GetAll() ([]model.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) GetById(id int64) (model.Category, error) {
	return s.repo.GetById(id)
}

func (s *CategoryService) GetByName(name string) (model.Category, error) {
	return s.repo.GetByName(name)
}
