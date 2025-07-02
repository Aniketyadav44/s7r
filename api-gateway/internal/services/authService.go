package services

import (
	"database/sql"
	"fmt"

	"github.com/Aniketyadav44/s7r/api-gateway/internal/models"
	"github.com/Aniketyadav44/s7r/api-gateway/pkg/utils"
)

type AuthService struct {
	db *sql.DB
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

// Register a new user by providing name, email, passowrd. It returns user and error
func (s *AuthService) Register(name, email, password string) (*models.User, error) {
	hashedPass, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	checkQuery := "SELECT 1 from users WHERE email = $1"
	var check int
	if err := s.db.QueryRow(checkQuery, email).Scan(&check); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if check == 1 {
		return nil, fmt.Errorf("user already exists with this email")
	}

	user := &models.User{
		Name:  name,
		Email: email,
	}
	query := "INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id"
	if err := s.db.QueryRow(query, name, email, hashedPass).Scan(&user.Id); err != nil {
		return nil, err
	}

	return user, nil
}

// Login a user by email and password. It returns user and error
func (s *AuthService) Login(email, password string) (*models.User, error) {
	query := "SELECT id, name, email, password, created_at FROM users where email = $1"

	user := &models.User{}
	var hashedPass string
	if err := s.db.QueryRow(query, email).Scan(&user.Id, &user.Name, &user.Email, &hashedPass, &user.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user doesn't exists with this email")
		}
		return nil, err
	}

	ok := utils.CheckPasswordHash(password, hashedPass)
	if !ok {
		return nil, fmt.Errorf("password didn't match")
	}

	return user, nil
}
