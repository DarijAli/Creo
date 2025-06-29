package unmarshal

import (
	"encoding/json"
	"templates/go/lib/register/src/models"
)

// Returns a User struct
func UnmarshalUser(jsonData []byte) (models.User, error) {
	var user models.User

	err := json.Unmarshal(jsonData, &user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
