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

	discounts, err := sdk.SetDiscountToPaymentMethodWithContext(context.Background(), cryptomus.SetDiscountToPaymentMethodRequest{
		Currency: "USD",
		Network:  "BTC",
		Discount: 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Discount successfully.")
	log.Printf("Discount: %#+v", discounts)
}
