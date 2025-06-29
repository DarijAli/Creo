package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"seed_login/src/db"
	"seed_login/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const batchSize = 5000

func batched(start, end, size int) <-chan []int {
	out := make(chan []int)
	go func() {
		defer close(out)
		if size < 1 {
			log.Fatal("batch size must be at least one")
		}
		for i := start; i <= end; i += size {
			to := min(i+size-1, end)
			batch := make([]int, 0, to-i+1)
			for j := i; j <= to; j++ {
				batch = append(batch, j)
			}
			out <- batch
		}
	}()
	return out
}

func randomUser(id int) map[string]interface{} {
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
	ctx := context.Background()
	collection := db.GetCollection()

	// Create indexes
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
		log.Fatalf("Failed to create indexes: %v", err)
	}

	for batch := range batched(1, seedCount, batchSize) {
		var docs []interface{}
		for _, id := range batch {
			docs = append(docs, randomUser(id))
		}

		_, err := collection.InsertMany(ctx, docs)
		if err != nil {
			log.Fatalf("Failed to insert batch: %v", err)
		}
	}

	log.Println("✅ Seeding complete")
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
