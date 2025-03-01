package main

import (
	"context"
	"log"
	"time"

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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	createResult, err := sdk.CreateRecurringPaymentWithContext(ctx, cryptomus.CreateRecurringPaymentRequest{
		Amount:   "15",
		Currency: "USDT",
		Name:     "Recurring payment",
		Period:   "monthly",
	})
	if err != nil {
		log.Fatal(err)
	}

	result, err := sdk.RecurringPaymentInformationWithContext(ctx, cryptomus.RecurringPaymentInformationRequest{
		UUID: createResult.Result.UUID,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Recurring payment information successfully.")
	log.Printf("Recurring payment information: %#+v", result)
}
