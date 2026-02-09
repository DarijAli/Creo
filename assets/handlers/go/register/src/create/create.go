package create

import (
	"context"
	"templates/go/lib/register/src/db"
	"templates/go/lib/register/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user *models.User) (string, error) {
	insertDoc, err := bson.Marshal(user)
	if err != nil {
		return "", err
	}

	var docMap bson.M
	err = bson.Unmarshal(insertDoc, &docMap)
	if err != nil {
		return "", err
	}

	// Insert document into collection
	result, err := db.RegisterCollection.InsertOne(context.Background(), docMap)
	if err != nil {
		return "", err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", nil
	}

	return oid.Hex(), nil
}
