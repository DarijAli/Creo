package models

import (
	"fmt"
	"templates/go/lib/user_create/src/hash"
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Username     string    `json:"username" validate:"required,min=3,max=64"`
	Email        string    `json:"email" validate:"required,min=3,max=64"`
	PasswordHash string    `json:"password" validate:"required,min=6,max=48"`
	CreatedAt    time.Time `json:"created_at"`
}

// NewUser validates input, hashes the password and returns a user ready for storage.
func NewUser(userData User) (*User, error) {
	validate := validator.New()
	if err := validate.Struct(userData); err != nil {
		return nil, err
	}

	passwordHash, err := hash.HashPassword(userData.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("hashing password: %w", err)
	}

	user := &User{
		Username:     userData.Username,
		Email:        userData.Email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now().UTC(),
	}

	return user, nil
}
