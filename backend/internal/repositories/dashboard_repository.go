package repositories

import (
	"car-rental-system/backend/internal/models"
	"database/sql"
)

type DashboardRepository struct {
	db *sql.DB
}

func NewDashboardRepository(
	db *sql.DB,
) *DashboardRepository {
	return &DashboardRepository{
		db: db,
	}
}

func (r *DashboardRepository) GetStats() (
	*models.DashboardStats,
	error,
) {

	var stats models.DashboardStats

	err := r.db.QueryRow(`
	SELECT
		(SELECT COUNT(*) FROM cars),
		(SELECT COUNT(*) FROM cars WHERE status = 'available'),
		(SELECT COUNT(*) FROM bookings WHERE status = 'new'),
		(SELECT COUNT(*) FROM bookings WHERE status IN ('approved','active')),
		COALESCE(
			(
				SELECT SUM(total_price)
				FROM bookings
				WHERE status IN ('completed','active')
			),
			0
		)
	`).Scan(
		&stats.CarsTotal,
		&stats.CarsAvailable,
		&stats.BookingsNew,
		&stats.BookingsActive,
		&stats.Revenue,
	)

	if err != nil {
		return nil, err
	}

	return &stats, nil
}
