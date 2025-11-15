package create

import (
	"context"
	"templates/go/lib/user_create/src/db"
	"templates/go/lib/user_create/src/models"
	"templates/go/lib/user_create/src/unmarshal"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Creates a new user in Database.
func CreateUser(jsonData []byte) (string, error) {
	// Unmarshal JSON data into a User struct
	userStruct, err := unmarshal.UnmarshalUser(jsonData)
	if err != nil {
		return "", err
	}

	db.InitMongoSafe()

	user, err := models.NewUser(userStruct)
	if err != nil {
		return "", err
	}

	insertResult, err := db.UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}

	id := insertResult.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}
