package services

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/0xDevvvvv/makerble/internal/models"
	"github.com/0xDevvvvv/makerble/internal/repositories"
	"github.com/0xDevvvvv/makerble/pkg/utils"
)

func ValidateUser(db *sql.DB, username, password string) (string, error) {
	userRepo := repositories.NewUserRepository(db)

	//get user by name
	user, err := userRepo.GetByUsername(username)

	if err != nil {
		return "", fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return "", errors.New("user not found")
	}
	//get password hash
	hash, err := userRepo.GetPassword(user.ID)
	if err != nil {
		return "", fmt.Errorf("failed to get password: %w", err)
	}
	if hash == "" {
		return "", errors.New("password not set")
	}
	//check password hash
	err = utils.CheckPassword(hash, password)
	if err != nil {
		return "", errors.New("invalid password")

	}

	//generate jwt token with username and role
	jwt, err := utils.GenerateToken(user.Username, user.Role)
	if err != nil {
		return "", errors.New("error generating jwt")
	}
	return jwt, nil
}

func CreateUser(db *sql.DB, user *models.UserCreate) (*models.UserCreate, error) {
	userRepo := repositories.NewUserRepository(db)
	hashedPassword, err := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	return userRepo.Create(user)
}
