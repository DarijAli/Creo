package models

import (
	"math/rand"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

func randomString(minLen, maxLen int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	n := minLen + rand.Intn(maxLen-minLen+1)
	var sb strings.Builder
	for range n {
		sb.WriteByte(letters[rand.Intn(len(letters))])
	}
	return sb.String()
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

type Address struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=64"`
	LastName  string `json:"last_name" validate:"required,min=2,max=64"`
	Street    string `json:"street" validate:"required,min=2,max=128"`
	Number    int    `json:"number" validate:"required,gt=0"`
	ZipCode   int    `json:"zip_code" validate:"required"`
	City      string `json:"city" validate:"required"`
	Country   string `json:"country" validate:"required"`
}

func NewAddress(addr *Address) *Address {
	if addr == nil {
		addr = &Address{}
	}
	if addr.FirstName == "" {
		addr.FirstName = randomString(2, 64)
	}
	if addr.LastName == "" {
		addr.LastName = randomString(2, 64)
	}
	if addr.Street == "" {
		addr.Street = randomString(2, 128)
	}
	if addr.Number == 0 {
		addr.Number = randomInt(1, 2000)
	}
	if addr.ZipCode == 0 {
		addr.ZipCode = randomInt(1000, 99999)
	}
	if addr.City == "" {
		addr.City = randomString(3, 64)
	}
	if addr.Country == "" {
		addr.Country = randomString(3, 64)
	}
	return addr
}

type Item struct {
	PriceInCents int    `json:"price_in_cents" validate:"required,gt=0"`
	Name         string `json:"name" validate:"required,min=1,max=128"`
}

func NewItem(item *Item) *Item {
	if item == nil {
		item = &Item{}
	}
	if item.PriceInCents == 0 {
		item.PriceInCents = randomInt(1, 1_000_000_000)
	}
	if item.Name == "" {
		item.Name = randomString(1, 128)
	}
	return item
}

type OrderItem struct {
	Item     Item `json:"item" validate:"required,dive"`
	Quantity int  `json:"quantity" validate:"required,gt=0"`
}

func NewOrderItem(orderItem *OrderItem) *OrderItem {
	if orderItem == nil {
		orderItem = &OrderItem{}
	}
	orderItem.Item = *NewItem(&orderItem.Item)
	if orderItem.Quantity == 0 {
		orderItem.Quantity = randomInt(1, 10_000)
	}
	return orderItem
}

type Invoice struct {
	Items           []OrderItem `json:"items" validate:"required,dive"`
	BillingAddress  Address     `json:"billing_address" validate:"required,dive"`
	ShippingAddress Address     `json:"shipping_address" validate:"required,dive"`
	UserID          string      `json:"user_id" validate:"required"`
	TaxRate         float64     `json:"tax_rate" validate:"gte=0"`
	IssuedAt        time.Time   `json:"issued_at"`
	ExtraInfo       string      `json:"extra_info"`
	Status          string      `json:"status" validate:"required,oneof=OPEN PAID"`
	InvoiceNumber   string      `json:"invoice_number" validate:"required,min=10,max=13"`
}

func NewInvoice(inv *Invoice) (*Invoice, error) {
	validate := validator.New()

	if inv == nil {
		inv = &Invoice{}
	}

	if len(inv.Items) == 0 {
		n := randomInt(1, 100)
		items := make([]OrderItem, n)
		for i := range n {
			items[i] = *NewOrderItem(nil)
		}
		inv.Items = items
	} else {
		for i := range inv.Items {
			inv.Items[i] = *NewOrderItem(&inv.Items[i])
		}
	}

	inv.BillingAddress = *NewAddress(&inv.BillingAddress)
	inv.ShippingAddress = *NewAddress(&inv.ShippingAddress)

	if inv.UserID == "" {
		inv.UserID = randomString(10, 24)
	}

	if inv.TaxRate == 0 {
		inv.TaxRate = 0.15
	}

	if inv.IssuedAt.IsZero() {
		inv.IssuedAt = time.Now().UTC()
	}

	if inv.Status == "" {
		inv.Status = "OPEN"
	}

	if inv.InvoiceNumber == "" {
		inv.InvoiceNumber = randomString(10, 13)
	}

	err := validate.Struct(inv)
	if err != nil {
		return nil, err
	}

	return inv, nil
}
