package services

import (
	"car-rental-system/backend/internal/models"
	"car-rental-system/backend/internal/repositories"
	"errors"
	"regexp"
)

type CustomerService struct {
	repo repositories.CustomerRepository
}

func NewCustomerService(
	repo repositories.CustomerRepository,
) *CustomerService {
	return &CustomerService{
		repo: repo,
	}
}

func (s *CustomerService) Create(
	customer *models.Customer,
) error {

	if err := validateCustomer(customer); err != nil {
		return err
	}

	return s.repo.Create(customer)
}

func (s *CustomerService) GetAll() ([]models.Customer, error) {
	return s.repo.GetAll()
}

func (s *CustomerService) GetByID(id int) (*models.Customer, error) {
	return s.repo.GetByID(id)
}

func (s *CustomerService) Update(
	customer *models.Customer,
) error {

	if err := validateCustomer(customer); err != nil {
		return err
	}

	return s.repo.Update(customer)
}

func (s *CustomerService) Delete(id int) error {
	return s.repo.Delete(id)
}

func validateCustomer(customer *models.Customer) error {

	phoneRegex := regexp.MustCompile(`^\+7\d{10}$`)

	if !phoneRegex.MatchString(customer.Phone) {
		return errors.New(
			"phone must be like +77001234567",
		)
	}

	iinRegex := regexp.MustCompile(`^\d{12}$`)

	if !iinRegex.MatchString(customer.IIN) {
		return errors.New(
			"IIN must contain exactly 12 digits",
		)
	}

	return nil
}
