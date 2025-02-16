package main

type Endpoint string

var (
	Host = "https://api.cryptomus.com"
)

var (
	// Payments
	CreateInvoice                 Endpoint = "/v1/payment"
	CreateStaticWallet            Endpoint = "/v1/wallet"
	GenerateQRCode                Endpoint = "/v1/wallet/qr"
	BlockStaticWallet             Endpoint = "/v1/wallet/block-address"
	RefundPaymentOnBlockedAddress Endpoint = "/v1/wallet/blocked-address-refund"
	PaymentInformation            Endpoint = "/v1/payment/info"
	Refund                        Endpoint = "/v1/payment/refund"
	ResendWebhook                 Endpoint = "/v1/payment/resend"
	TestingWebhook                Endpoint = "/v1/test-webhook/payment"
	PaymentListOfServices         Endpoint = "/v1/payment/services"
	PaymentHistory                Endpoint = "/v1/payment/list"

	// Payouts
	CreatePayout             Endpoint = "/v1/payout"
	PayoutInformation        Endpoint = "/v1/payout/info"
	PayoutHistory            Endpoint = "/v1/payout/list"
	PayoutListOfServices     Endpoint = "/v1/payout/services"
	TransferToPersonalWallet Endpoint = "/v1/transfer/to-personal"
	TransferToBusinessWallet Endpoint = "/v1/transfer/to-business"

	// Recurring Payments
	CreateRecurringPayment      Endpoint = "/v1/recurrence/create"
	RecurringPaymentInformation Endpoint = "/v1/recurrence/info"
	ListRecurringPayments       Endpoint = "/v1/recurrence/list"
	CancelRecurringPayment      Endpoint = "/v1/recurrence/cancel"

	// Exchange Rate
	ExchangeRateList Endpoint = "/v1/exchange-rate/{currency}/list"

	// Discount Payment
	ListOfDiscounts            Endpoint = "/v1/payment/discount/list"
	SetDiscountToPaymentMethod Endpoint = "/v1/payment/discount/set"

	// Balance
	Balance Endpoint = "/v1/balance"
)
