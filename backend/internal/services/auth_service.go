package services

import (
	"os"
	"time"

	"car-rental-system/backend/internal/models"
	"car-rental-system/backend/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(
	userRepo *repositories.UserRepository,
) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Register(
	fullName,
	email,
	password string,
) error {

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := models.User{
		FullName:     fullName,
		Email:        email,
		PasswordHash: string(hash),
		Role:         "admin",
	}

	return s.userRepo.Create(&user)
}

func (s *AuthService) Login(
	email,
	password string,
) (string, error) {

	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": user.ID,
			"role":    user.Role,
			"exp":     time.Now().Add(72 * time.Hour).Unix(),
		},
	)

	return token.SignedString(
		[]byte(os.Getenv("JWT_SECRET")),
	)
}
