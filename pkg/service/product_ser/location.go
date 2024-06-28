package product_ser

import (
	"github.com/sirupsen/logrus"
	"ims-product-api/model"
	"ims-product-api/pkg/repository"
)

type LocationService struct {
	repo repository.Location
}

func NewLocationService(repo repository.Location) *LocationService {
	return &LocationService{
		repo: repo,
	}
}

func (s *LocationService) Create(location model.Location) (int64, error) {
	return s.repo.Create(location)
}

func (s *LocationService) Update(id int64, location model.UpdateLocation) error {
	if err := model.Verify(location); err != nil {
		logrus.Error(err)
		return err
	}

	return s.repo.Update(id, location)
}

func (s *LocationService) Delete(id int64) error {
	return s.repo.Delete(id)
}

func (s *LocationService) GetAll() ([]model.Location, error) {
	return s.repo.GetAll()
}

func (s *LocationService) GetById(id int64) (model.Location, error) {
	return s.repo.GetById(id)
}

func (s *LocationService) GetByRowAndPlace(row, place string) (model.Location, error) {
	return s.repo.GetByRowAndPlace(row, place)
}

func (s *LocationService) GetByRow(row string) ([]model.Location, error) {
	return s.repo.GetListByRow(row)
}
