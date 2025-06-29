package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"seed_register/src/db"
	"seed_register/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const batchSize = 5000

func batched(ids []int, n int) <-chan []int {
	out := make(chan []int)
	go func() {
		defer close(out)
		if n < 1 {
			log.Fatal("batch size must be at least one")
		}
		for i := 0; i < len(ids); i += n {
			end := min(i+n, len(ids))
			out <- ids[i:end]
		}
	}()
	return out
}

func randomUser(id int) map[string]any {
	user := models.NewRandomUser()
	return map[string]interface{}{
		"_id":        id,
		"username":   user.Username,
		"email":      user.Email,
		"password":   user.PasswordHash,
		"created_at": user.CreatedAt,
	}
}

func seedDatabase(seedCount int) error {
	collection := db.GetCollection()
	ctx := context.Background()

	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "username", Value: 1}},
		},
		{
			Keys: bson.D{{Key: "email", Value: 1}},
		},
	}
	_, err := collection.Indexes().CreateMany(ctx, indexes)
	if err != nil {
		return err
	}

	ids := make([]int, seedCount)
	for i := 0; i < seedCount; i++ {
		ids[i] = i + 1
	}

	for batch := range batched(ids, batchSize) {
		var docs []any
		for _, id := range batch {
			docs = append(docs, randomUser(id))
		}
		insertResult, err := collection.InsertMany(ctx, docs)
		if err != nil {
			return err
		}
		if insertResult == nil || len(insertResult.InsertedIDs) == 0 {
			return fmt.Errorf("failed to insert batch: no inserted IDs returned")
		}
	}

	return nil
}

func main() {
	seedCountStr := os.Getenv("MG_SEED_COUNT")
	if seedCountStr == "" {
		log.Fatal("MG_SEED_COUNT env var not set")
	}

	seedCount, err := strconv.Atoi(seedCountStr)
	if err != nil {
		log.Fatalf("Invalid MG_SEED_COUNT: %v", err)
	}

	if err := seedDatabase(seedCount); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	log.Println("✅ Database seeding completed successfully")
}
