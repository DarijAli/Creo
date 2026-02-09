package create

import (
	"context"

	"templates/go/lib/invoice_create/src/db"
	"templates/go/lib/invoice_create/src/models"
	"templates/go/lib/invoice_create/src/unmarshal"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateInvoice(jsonData []byte) (string, error) {
	invoiceStruct, err := unmarshal.UnmarshalInvoice(jsonData)
	if err != nil {
		return "", err
	}

	db.InitMongoSafe()

	invoice, err := models.NewInvoice(&invoiceStruct)
	if err != nil {
		return "", err
	}

	insertResult, err := db.InvoiceCollection.InsertOne(context.Background(), invoice)
	if err != nil {
		return "", err
	}

	id := insertResult.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}
