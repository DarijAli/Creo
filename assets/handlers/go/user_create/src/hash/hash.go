package hash

import (
	"fmt"

	argon2 "github.com/alexedwards/argon2id"
)

const (
	TIME_COST   = 1
	MEMORY_COST = 6144
	SALT_LEN    = 16
	KEY_LEN     = 32
)

// HashPassword hashes a password using Argon2id and returns the hash or an error.
// It no longer panics on failure so callers can handle errors gracefully.
func HashPassword(password string) (string, error) {
	params := &argon2.Params{
		Memory:      MEMORY_COST,
		Iterations:  TIME_COST,
		Parallelism: 2,
		SaltLength:  SALT_LEN,
		KeyLength:   KEY_LEN,
	}

	hashPassword, err := argon2.CreateHash(password, params)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return hashPassword, nil
}
