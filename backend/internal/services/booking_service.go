package services

import (
	"car-rental-system/backend/internal/models"
	"car-rental-system/backend/internal/repositories"
	"errors"
)

type BookingService struct {
	repo repositories.BookingRepository
}

func NewBookingService(
	repo repositories.BookingRepository,
) *BookingService {
	return &BookingService{
		repo: repo,
	}
}

func (s *BookingService) Create(
	booking *models.Booking,
) error {

	available, err := s.repo.IsCarAvailable(
		booking.CarID,
		booking.StartDate,
		booking.EndDate,
	)

	if err != nil {
		return err
	}

	if !available {
		return errors.New(
			"car already booked for selected dates",
		)
	}

	return s.repo.Create(booking)
}
func (s *BookingService) GetAll() ([]models.Booking, error) {
	return s.repo.GetAll()
}

func (s *BookingService) GetByID(id int) (*models.Booking, error) {
	return s.repo.GetByID(id)
}
func (s *BookingService) UpdateStatus(
	id int,
	status string,
) error {
	return s.repo.UpdateStatus(id, status)
}
