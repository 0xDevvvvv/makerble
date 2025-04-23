package repositories

import (
	"database/sql"
	"fmt"

	"github.com/0xDevvvvv/makerble/internal/models"
)

type UserRepository interface {
	GetById(id int) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetPassword(id int) (string, error)
	Create(user *models.UserCreate) (*models.UserCreate, error)
	// Update(user *models.User) error
	// Delete(user *models.User) error
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) GetByUsername(username string) (*models.User, error) {
	var userDetails models.User

	query := `
		SELECT id, username, role, created_at
		FROM users
		WHERE username = $1;
	`

	row := r.db.QueryRow(query, username)

	err := row.Scan(&userDetails.ID, &userDetails.Username, &userDetails.Role, &userDetails.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // not found
		}
		return nil, fmt.Errorf("error Fetching Record (By username) : %w", err)
	}

	return &userDetails, nil

}

func (r *userRepo) GetById(id int) (*models.User, error) {
	var userDetails models.User
	query := `
		SELECT id, username, role, created_at
		FROM users
		WHERE id = $1;
	`

	row := r.db.QueryRow(query, id)

	err := row.Scan(&userDetails.ID, &userDetails.Username, &userDetails.Role, &userDetails.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // not found
		}
		return nil, fmt.Errorf("error Fetching Record (By username) : %w", err)
	}

	return &userDetails, nil

}
func (r *userRepo) GetPassword(id int) (string, error) {
	var hash string
	query := `
		SELECT password_hash FROM user_passwords
		WHERE id = $1;
	`

	row := r.db.QueryRow(query, id)
	err := row.Scan(&hash)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil // not found
		}
		return "", fmt.Errorf("error Fetching Hash (By ID) : %w", err)
	}
	return hash, nil
}

func (r *userRepo) Create(user *models.UserCreate) (*models.UserCreate, error) {

	query := `
		INSERT INTO users (username, role)
		VALUES ($1, $2)
		RETURNING id, created_at;
	`

	err := r.db.QueryRow(query, user.Username, user.Role).
		Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("error creating user (CreateUser): %w", err)
	}

	query = `
		INSERT INTO user_passwords (id, password_hash)
		VALUES ($1, $2)
		RETURNING id, password_hash;
	`
	err = r.db.QueryRow(query, user.ID, user.Password).
		Scan(&user.ID, &user.Password)

	if err != nil {
		return nil, fmt.Errorf("error creating user (CreateUser): %w", err)
	}

	fmt.Printf("User Created : %v %v %v %v %v\n", user.Username, user.Role, user.ID, user.CreatedAt, user.Password)
	return user, nil
}

// func (r *userRepo) Update(user *models.User) error {
// 	return nil
// }

// func (r *userRepo) Delete(user *models.User) error {
// 	return nil
// }
