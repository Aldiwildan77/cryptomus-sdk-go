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

	log.Println("Creating recurring payment...")

	result, err := sdk.CreateRecurringPaymentWithContext(context.Background(), cryptomus.CreateRecurringPaymentRequest{
		Amount:   "15",
		Currency: "USDT",
		Name:     "Recurring payment",
		Period:   "monthly",
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Recurring payment successfully.")
	log.Printf("Recurring payment: %#+v", result)
}
