package main

import (
	"context"
	"fmt"
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

	resp, err := sdk.ExchangeRateListWithContext(context.Background(), "USD")
	if err != nil {
		log.Fatal(err)
	}

	for _, rate := range resp.Result {
		fmt.Printf("From: %s, To: %s, Course: %s\n", rate.From, rate.To, rate.Course)
	}
}
