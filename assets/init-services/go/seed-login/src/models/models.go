package models

import (
	"math/rand"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

func randomString(minLen, maxLen int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	n := minLen + rand.Intn(maxLen-minLen+1)
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteByte(letters[rand.Intn(len(letters))])
	}
	return sb.String()
}

type User struct {
	Username     string    `json:"username" validate:"required,min=3,max=64"`
	Email        string    `json:"email" validate:"required,email,min=3,max=64"`
	PasswordHash string    `json:"password" validate:"required,min=32,max=128"`
	CreatedAt    time.Time `json:"created_at"`
}

func NewRandomUser() *User {
	username := randomString(3, 64)
	email := randomString(3, 64) + "@example.com"
	password := randomString(97, 97)

	user := &User{
		Username:     username,
		Email:        email,
		PasswordHash: password,
		CreatedAt:    time.Now().UTC(),
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		panic("Random user generation failed validation: " + err.Error())
	}

	return user
}
