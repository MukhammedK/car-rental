package repositories

import (
	"car-rental-system/backend/internal/models"
	"database/sql"
)

type CustomerRepository interface {
	Create(customer *models.Customer) error
	GetAll() ([]models.Customer, error)
	GetByID(id int) (*models.Customer, error)
	Update(customer *models.Customer) error
	Delete(id int) error
}

type PostgresCustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *PostgresCustomerRepository {
	return &PostgresCustomerRepository{
		db: db,
	}
}

func (r *PostgresCustomerRepository) Create(customer *models.Customer) error {

	query := `
	INSERT INTO customers(
		full_name,
		phone,
		email,
		iin
	)
	VALUES($1,$2,$3,$4)
	RETURNING id
	`

	return r.db.QueryRow(
		query,
		customer.FullName,
		customer.Phone,
		customer.Email,
		customer.IIN,
	).Scan(&customer.ID)
}

func (r *PostgresCustomerRepository) GetAll() ([]models.Customer, error) {

	rows, err := r.db.Query(`
		SELECT
			id,
			full_name,
			phone,
			email,
			iin,
			created_at
		FROM customers
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var customers []models.Customer

	for rows.Next() {

		var customer models.Customer

		err := rows.Scan(
			&customer.ID,
			&customer.FullName,
			&customer.Phone,
			&customer.Email,
			&customer.IIN,
			&customer.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

func (r *PostgresCustomerRepository) GetByID(id int) (*models.Customer, error) {

	var customer models.Customer

	err := r.db.QueryRow(`
		SELECT
			id,
			full_name,
			phone,
			email,
			iin,
			created_at
		FROM customers
		WHERE id = $1
	`, id).Scan(
		&customer.ID,
		&customer.FullName,
		&customer.Phone,
		&customer.Email,
		&customer.IIN,
		&customer.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *PostgresCustomerRepository) Update(customer *models.Customer) error {

	_, err := r.db.Exec(`
		UPDATE customers
		SET
			full_name = $1,
			phone = $2,
			email = $3,
			iin = $4
		WHERE id = $5
	`,
		customer.FullName,
		customer.Phone,
		customer.Email,
		customer.IIN,
		customer.ID,
	)

	return err
}

func (r *PostgresCustomerRepository) Delete(id int) error {

	_, err := r.db.Exec(`
		DELETE FROM customers
		WHERE id = $1
	`, id)

	return err
}
