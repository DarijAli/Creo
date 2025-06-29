package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"seed_invoice/src/db"
	"seed_invoice/src/models"
)

const batchSize = 50000

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

func seedDatabase(seedCount int) error {
	collection := db.GetCollection()
	ctx := context.Background()
	for batch := range batched(1, seedCount, batchSize) {
		var docs []interface{}
		for _, id := range batch {
			inv, err := models.NewInvoice(nil)
			if err != nil {
				return fmt.Errorf("failed to create invoice: %w", err)
			}

			doc := map[string]interface{}{
				"_id":              id,
				"items":            inv.Items,
				"billing_address":  inv.BillingAddress,
				"shipping_address": inv.ShippingAddress,
				"user_id":          inv.UserID,
				"tax_rate":         inv.TaxRate,
				"issued_at":        inv.IssuedAt,
				"extra_info":       inv.ExtraInfo,
				"status":           inv.Status,
				"invoice_number":   inv.InvoiceNumber,
			}

			docs = append(docs, doc)
		}

		_, err := collection.InsertMany(ctx, docs)
		if err != nil {
			return fmt.Errorf("failed to insert batch: %w", err)
		}
		fmt.Printf("Inserted batch of %d invoices\n", len(docs))
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
