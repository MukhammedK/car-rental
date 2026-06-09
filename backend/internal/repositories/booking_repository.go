package repositories

import (
	"car-rental-system/backend/internal/models"
	"database/sql"
	"time"
)

type BookingRepository interface {
	Create(booking *models.Booking) error

	GetAll() ([]models.Booking, error)

	GetByID(id int) (*models.Booking, error)

	IsCarAvailable(
		carID int,
		startDate time.Time,
		endDate time.Time,
	) (bool, error)
	UpdateStatus(
		id int,
		status string,
	) error
}

type PostgresBookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(
	db *sql.DB,
) *PostgresBookingRepository {
	return &PostgresBookingRepository{
		db: db,
	}
}

func (r *PostgresBookingRepository) IsCarAvailable(
	carID int,
	startDate time.Time,
	endDate time.Time,
) (bool, error) {

	var count int

	query := `
	SELECT COUNT(*)
	FROM bookings
	WHERE car_id = $1
	AND status IN ('new', 'approved')
	AND (
		start_date <= $3
		AND end_date >= $2
	)
	`

	err := r.db.QueryRow(
		query,
		carID,
		startDate,
		endDate,
	).Scan(&count)

	if err != nil {
		return false, err
	}

	return count == 0, nil
}
func (r *PostgresBookingRepository) Create(
	booking *models.Booking,
) error {

	query := `
	INSERT INTO bookings(
		customer_id,
		car_id,
		start_date,
		end_date,
		total_price,
		status,
		source,
		comment
	)
	VALUES($1,$2,$3,$4,$5,$6,$7,$8)
	RETURNING id
	`

	return r.db.QueryRow(
		query,
		booking.CustomerID,
		booking.CarID,
		booking.StartDate,
		booking.EndDate,
		booking.TotalPrice,
		booking.Status,
		booking.Source,
		booking.Comment,
	).Scan(&booking.ID)
}

func (r *PostgresBookingRepository) GetAll() ([]models.Booking, error) {

	rows, err := r.db.Query(`
		SELECT
			id,
			customer_id,
			car_id,
			start_date,
			end_date,
			total_price,
			status,
			source,
			comment,
			created_at,
			updated_at
		FROM bookings
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var bookings []models.Booking

	for rows.Next() {

		var booking models.Booking

		err := rows.Scan(
			&booking.ID,
			&booking.CustomerID,
			&booking.CarID,
			&booking.StartDate,
			&booking.EndDate,
			&booking.TotalPrice,
			&booking.Status,
			&booking.Source,
			&booking.Comment,
			&booking.CreatedAt,
			&booking.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func (r *PostgresBookingRepository) GetByID(id int) (*models.Booking, error) {

	var booking models.Booking

	err := r.db.QueryRow(`
		SELECT
			id,
			customer_id,
			car_id,
			start_date,
			end_date,
			total_price,
			status,
			source,
			comment,
			created_at,
			updated_at
		FROM bookings
		WHERE id = $1
	`, id).Scan(
		&booking.ID,
		&booking.CustomerID,
		&booking.CarID,
		&booking.StartDate,
		&booking.EndDate,
		&booking.TotalPrice,
		&booking.Status,
		&booking.Source,
		&booking.Comment,
		&booking.CreatedAt,
		&booking.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &booking, nil
}

func (r *PostgresBookingRepository) UpdateStatus(
	id int,
	status string,
) error {

	_, err := r.db.Exec(`
		UPDATE bookings
		SET
			status = $1,
			updated_at = NOW()
		WHERE id = $2
	`,
		status,
		id,
	)

	return err
}
