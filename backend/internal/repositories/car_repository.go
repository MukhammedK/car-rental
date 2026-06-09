package repositories

import (
	"car-rental-system/backend/internal/models"
	"database/sql"
	"fmt"
)

type CarRepository interface {
	Create(car *models.Car) error
	GetAll() ([]models.Car, error)
	GetByID(id int) (*models.Car, error)
	Update(car *models.Car) error
	Delete(id int) error
}
type PostgresCarRepository struct {
	db *sql.DB
}

func NewCarRepository(db *sql.DB) *PostgresCarRepository {
	return &PostgresCarRepository{
		db: db,
	}
}

func (r *PostgresCarRepository) Create(car *models.Car) error {

	query := `
	INSERT INTO cars (
		brand,
		model,
		year,
		license_plate,
		color,
		transmission,
		fuel_type,
		daily_price,
		status,
		description
	)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
	RETURNING id
	`

	err := r.db.QueryRow(
		query,
		car.Brand,
		car.Model,
		car.Year,
		car.LicensePlate,
		car.Color,
		car.Transmission,
		car.FuelType,
		car.DailyPrice,
		car.Status,
		car.Description,
	).Scan(&car.ID)

	if err != nil {
		fmt.Println("CREATE ERROR:", err)
		return err
	}

	return nil
}

func (r *PostgresCarRepository) GetAll() ([]models.Car, error) {

	query := `
	SELECT
		id,
		brand,
		model,
		year,
		license_plate,
		color,
		transmission,
		fuel_type,
		daily_price,
		status,
		description,
		created_at
	FROM cars
	ORDER BY id DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []models.Car

	for rows.Next() {

		var car models.Car

		err := rows.Scan(
			&car.ID,
			&car.Brand,
			&car.Model,
			&car.Year,
			&car.LicensePlate,
			&car.Color,
			&car.Transmission,
			&car.FuelType,
			&car.DailyPrice,
			&car.Status,
			&car.Description,
			&car.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		cars = append(cars, car)
	}

	return cars, nil
}

func (r *PostgresCarRepository) GetByID(id int) (*models.Car, error) {

	query := `
	SELECT
		id,
		brand,
		model,
		year,
		license_plate,
		color,
		transmission,
		fuel_type,
		daily_price,
		status,
		description,
		created_at
	FROM cars
	WHERE id = $1
	`

	var car models.Car

	err := r.db.QueryRow(query, id).Scan(
		&car.ID,
		&car.Brand,
		&car.Model,
		&car.Year,
		&car.LicensePlate,
		&car.Color,
		&car.Transmission,
		&car.FuelType,
		&car.DailyPrice,
		&car.Status,
		&car.Description,
		&car.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &car, nil
}
func (r *PostgresCarRepository) Update(car *models.Car) error {

	query := `
	UPDATE cars
	SET
		brand = $1,
		model = $2,
		year = $3,
		license_plate = $4,
		color = $5,
		transmission = $6,
		fuel_type = $7,
		daily_price = $8,
		status = $9,
		description = $10
	WHERE id = $11
	`

	_, err := r.db.Exec(
		query,
		car.Brand,
		car.Model,
		car.Year,
		car.LicensePlate,
		car.Color,
		car.Transmission,
		car.FuelType,
		car.DailyPrice,
		car.Status,
		car.Description,
		car.ID,
	)

	return err
}

func (r *PostgresCarRepository) Delete(id int) error {

	query := `
	DELETE FROM cars
	WHERE id = $1
	`

	_, err := r.db.Exec(query, id)

	return err
}
