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

	result, err := sdk.ListRecurringPaymentsWithContext(context.Background(), cryptomus.ListRecurringPaymentsRequest{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Recurring payments successfully.")
	log.Printf("Recurring payments: %#+v", result)
}
