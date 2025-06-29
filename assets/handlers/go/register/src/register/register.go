package register

import (
	"encoding/json"
	"templates/go/lib/register/src/create"
	"templates/go/lib/register/src/models"
	"templates/go/lib/register/src/read"
)

func RegisterUser(jsonData []byte) (string, error) {
	var user models.User

	if err := json.Unmarshal(jsonData, &user); err != nil {
		return "", err
	}

	existingByEmail, err := read.ReadUserByEmail(user.Email)
	if err != nil {
		return "", err
	}
	if existingByEmail != nil {
		return "", nil
	}

	existingByUsername, err := read.ReadUserByUsername(user.Username)
	if err != nil {
		return "", err
	}
	if existingByUsername != nil {
		return "", nil
	}

	userID, err := create.CreateUser(&user)
	if err != nil {
		return "", err
	}

	return userID, nil
}
