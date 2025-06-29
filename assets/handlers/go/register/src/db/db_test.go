package db

import (
	"context"
	log "fmt"
	"testing"
	"time"
)

func TestInitMongo(t *testing.T) {
	InitMongo()

	if RegisterDb == nil {
		t.Errorf("Expected a non-nil database connection, got nil")
	}

	if RegisterCollection == nil {
		t.Errorf("Expected a non-nil collection, got nil")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := RegisterCollection.CountDocuments(ctx, map[string]any{})
	if err != nil {
		t.Errorf("Failed to count documents: %v", err)
	}

	log.Printf("📦 Total documents in 'register_collection': %d\n", count)

	if count != 0 {
		t.Errorf("Expected 0 documents, but got %d", count)
	}
}
