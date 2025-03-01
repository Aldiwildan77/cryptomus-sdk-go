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
	ctx := context.Background()

	sdk := cryptomus.New(
		cryptomus.WithMerchant(merchantID),
		cryptomus.WithPaymentToken(paymentToken),
	)

	balances, err := sdk.BalanceWithContext(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Balance successfully.")
	log.Printf("Balances: %#+v", balances)
}
