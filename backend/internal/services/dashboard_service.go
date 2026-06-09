package services

import (
	"car-rental-system/backend/internal/models"
	"car-rental-system/backend/internal/repositories"
)

type DashboardService struct {
	repo *repositories.DashboardRepository
}

func NewDashboardService(
	repo *repositories.DashboardRepository,
) *DashboardService {
	return &DashboardService{
		repo: repo,
	}
}

func (s *DashboardService) GetStats() (
	*models.DashboardStats,
	error,
) {
	return s.repo.GetStats()
}
