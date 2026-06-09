package services

import (
	"car-rental-system/backend/internal/models"
	"car-rental-system/backend/internal/repositories"
)

type CarService struct {
	repo repositories.CarRepository
}

func NewCarService(repo repositories.CarRepository) *CarService {
	return &CarService{
		repo: repo,
	}
}

func (s *CarService) Create(car *models.Car) error {
	return s.repo.Create(car)
}

func (s *CarService) GetAll() ([]models.Car, error) {
	return s.repo.GetAll()
}

func (s *CarService) GetByID(id int) (*models.Car, error) {
	return s.repo.GetByID(id)
}

func (s *CarService) Update(car *models.Car) error {
	return s.repo.Update(car)
}

func (s *CarService) Delete(id int) error {
	return s.repo.Delete(id)
}
