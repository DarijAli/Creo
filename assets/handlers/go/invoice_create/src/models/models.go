package models

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type Address struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=64"`
	LastName  string `json:"last_name" validate:"required,min=2,max=64"`
	Street    string `json:"street" validate:"required,min=2,max=128"`
	Number    int    `json:"number" validate:"required,gt=0"`
	ZipCode   int    `json:"zip_code" validate:"required"`
	City      string `json:"city" validate:"required"`
	Country   string `json:"country" validate:"required"`
}

type Item struct {
	PriceInCents int    `json:"price_in_cents" validate:"required,gt=0"`
	Name         string `json:"name" validate:"required,min=1,max=128"`
}

type OrderItem struct {
	Item     Item `json:"item" validate:"required"`
	Quantity int  `json:"quantity" validate:"required,gt=0"`
}

type Invoice struct {
	Items           []OrderItem `json:"items" validate:"required"`
	BillingAddress  Address     `json:"billing_address" validate:"required"`
	ShippingAddress Address     `json:"shipping_address" validate:"required"`
	UserID          string      `json:"user_id" validate:"required"`
	TaxRate         float64     `json:"tax_rate" validate:"gte=0"`
	IssuedAt        time.Time   `json:"issued_at"`
	ExtraInfo       string      `json:"extra_info"`
	Status          string      `json:"status" validate:"required,oneof=OPEN PAID"`
	InvoiceNumber   string      `json:"invoice_number" validate:"required,min=10,max=13"`
}

// Constructor function to create and validate new invoice.
func NewInvoice(inv *Invoice) (*Invoice, error) {
	if inv == nil {
		return nil, fmt.Errorf("invoice cannot be nil")
	}

	// Default time
	if inv.IssuedAt.IsZero() {
		inv.IssuedAt = time.Now().UTC()
	}

	// Default status
	if inv.Status == "" {
		inv.Status = "OPEN"
	}

	validate := validator.New()
	if err := validate.Struct(inv); err != nil {
		return nil, err
	}

	return inv, nil
}
