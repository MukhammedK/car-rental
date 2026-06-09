package repositories

import (
	"car-rental-system/backend/internal/models"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {

	query := `
	INSERT INTO users (
		full_name,
		email,
		password_hash,
		role
	)
	VALUES ($1,$2,$3,$4)
	RETURNING id
	`

	return r.db.QueryRow(
		query,
		user.FullName,
		user.Email,
		user.PasswordHash,
		user.Role,
	).Scan(&user.ID)
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {

	query := `
	SELECT
		id,
		full_name,
		email,
		password_hash,
		role,
		created_at
	FROM users
	WHERE email = $1
	`

	var user models.User

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
