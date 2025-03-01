# Cryptomus SDK Go

Cryptomus SDK Go is a Go library for interacting with the Cryptomus API. You can see the documentation for the Cryptomus API at [https://doc.cryptomus.com/](https://doc.cryptomus.com/).

## Features

- Payments Integration (Create Invoice, Get Invoice, Get Invoices, Cancel Invoice, etc.)
- Payout Integration (Create Payout, Get Payout, Get Payouts, Cancel Payout, etc.)
- Recurring Payment Integration (Create Recurring Payment, Get Recurring Payment, Get Recurring, Cancel Recurring Payment, etc.)
- Webhooks Integration (Verify Webhook, Resend Webhook)
- Static Wallet
- Support Balance Information
- Support Exchange Rate Information

## Installation

To install the Cryptomus SDK Go, you need to run the following command:

[![Go Reference](https://pkg.go.dev/badge/github.com/Aldiwildan77/cryptomus-sdk-go.svg)](https://pkg.go.dev/github.com/Aldiwildan77/cryptomus-sdk-go)

```bash
go get github.com/Aldiwildan77/cryptomus-sdk-go
```

## Usage

Here is an example of how to use the Cryptomus SDK Go, you can see the full example in the [examples](examples) directory.

### Create Invoice

```go
package main

import (
 "log"

 cryptomus "github.com/Aldiwildan77/cryptomus-sdk-go"
)

func main() {
 sdk := cryptomus.New(
  cryptomus.WithMerchant("fill your merchant id"),
  cryptomus.WithPaymentToken("fill your payment token"),
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
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

[Muhammad Wildan Aldiansyah](https://aldiwildan.me)
