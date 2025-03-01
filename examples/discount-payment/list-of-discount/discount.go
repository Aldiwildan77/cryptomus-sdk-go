package main

import (
	"context"
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

	discounts, err := sdk.ListOfDiscountWithContext(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Discounts successfully.")
	log.Printf("Discounts: %#+v", discounts)
}
