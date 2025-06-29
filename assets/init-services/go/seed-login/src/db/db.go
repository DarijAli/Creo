package db

import (
	"context"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCollection() *mongo.Collection {
	host := os.Getenv("DB_MONGO_HOST")
	portStr := os.Getenv("DB_MONGO_PORT")
	user := os.Getenv("DB_MONGO_USER")
	password := os.Getenv("DB_MONGO_PASSWORD")

	port, _ := strconv.Atoi(portStr)

	clientOptions := options.Client().
		SetHosts([]string{host + ":" + strconv.Itoa(port)}).
		SetAuth(options.Credential{
			Username: user,
			Password: password,
		})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	db := client.Database("login_db")
	return db.Collection("login_collection")
}
