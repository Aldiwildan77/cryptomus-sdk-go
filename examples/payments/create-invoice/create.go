package main

import (
	"log"

	cryptomus "github.com/Aldiwildan77/cryptomus-sdk-go"
)

const (
	merchantID   = ""
	paymentToken = ""
)

func main() {
	sdk := cryptomus.New(
		cryptomus.WithMerchant(merchantID),
		cryptomus.WithPaymentToken(paymentToken),
	)

	result, err := sdk.CreateInvoice(&cryptomus.CreateInvoiceRequest{
		Amount:   "15",
		Currency: "USD",
		OrderID:  "123456",
		Lifetime: 3600,
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Create invoice successfully.")
	log.Printf("Create invoice: %#+v", result)
}
