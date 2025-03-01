package cryptomus

import (
	"context"
	"time"
)

type Currency struct {
	Currency string `json:"currency"`
	Network  string `json:"network"`
}

type Currencies []Currency

type CreateInvoiceRequest struct {
	Amount                 string     `json:"amount"`
	Currency               string     `json:"currency"`
	OrderID                string     `json:"order_id"`
	Network                string     `json:"network"`
	URLReturn              string     `json:"url_return"`
	URLSuccess             string     `json:"url_success"`
	URLCallback            string     `json:"url_callback"`
	IsPaymentMultiple      bool       `json:"is_payment_multiple"`
	Lifetime               int        `json:"lifetime"`
	ToCurrency             string     `json:"to_currency"`
	Subtract               int        `json:"subtract"`
	AccuracyPaymentPercent int        `json:"accuracy_payment_percent"`
	AdditionalData         string     `json:"additional_data"`
	Currencies             Currencies `json:"currencies"`
	ExceptCurrencies       Currencies `json:"except_currencies"`
	CourseSource           string     `json:"course_source"`
	FromReferralCode       string     `json:"from_referral_code"`
	DiscountPercent        int        `json:"discount_percent"`
	IsRefresh              bool       `json:"is_refresh"`
}

type CreateInvoiceData struct {
	UUID            string    `json:"uuid"`
	OrderID         string    `json:"order_id"`
	Amount          string    `json:"amount"`
	PaymentAmount   int       `json:"payment_amount"`
	PayerAmount     int       `json:"payer_amount"`
	DiscountPercent int       `json:"discount_percent"`
	Discount        string    `json:"discount"`
	PayerCurrency   string    `json:"payer_currency"`
	Currency        string    `json:"currency"`
	MerchantAmount  int       `json:"merchant_amount"`
	Network         string    `json:"network"`
	Address         string    `json:"address"`
	From            string    `json:"from"`
	Txid            string    `json:"txid"`
	PaymentStatus   string    `json:"payment_status"`
	URL             string    `json:"url"`
	ExpiredAt       int       `json:"expired_at"`
	Status          string    `json:"status"`
	IsFinal         bool      `json:"is_final"`
	AdditionalData  string    `json:"additional_data"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CreateInvoiceResponse struct {
	*HTTPResponse
	Result CreateInvoiceData `json:"result"`
}

// CreateInvoice creates an invoice.
//
// Details: https://doc.cryptomus.com/business/payments/creating-invoice
//
// The invoice will have a specific cryptocurrency and address at the time of creation only if currency or to_currency parameter is a cryptocurrency and the network parameter is passed (or a cryptocurrency has only one network, for example BTC).
//
// Example:
//
//	result, err := sdk.CreateInvoice(&CreateInvoiceRequest{
//		Amount:     "15",
//		Currency:   "USD",
//		OrderID: 	"9d7f7b9b-3b7b-4b7b-9b7b-7b9d7f7b9b7b",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) CreateInvoice(payload *CreateInvoiceRequest) (*CreateInvoiceResponse, error) {
	return sdk.CreateInvoiceWithContext(context.Background(), payload)
}

// CreateInvoiceWithContext creates an invoice.
//
// Details: https://doc.cryptomus.com/business/payments/creating-invoice
//
// The invoice will have a specific cryptocurrency and address at the time of creation only if currency or to_currency parameter is a cryptocurrency and the network parameter is passed (or a cryptocurrency has only one network, for example BTC).
//
// Example:
//
//	result, err := sdk.CreateInvoiceWithContext(ctx, &CreateInvoiceRequest{
//		Amount:     "15",
//		Currency:   "USD",
//		OrderID: 	"9d7f7b9b-3b7b-4b7b-9b7b-7b9d7f7b9b7b",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) CreateInvoiceWithContext(ctx context.Context, payload *CreateInvoiceRequest) (*CreateInvoiceResponse, error) {
	var result CreateInvoiceResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, payload)).
		SetBody(payload).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(CreateInvoiceEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type CreateStaticWalletRequest struct {
	Currency         string `json:"currency"`
	Network          string `json:"network"`
	OrderID          string `json:"order_id"`
	URLCallback      string `json:"url_callback,omitempty"`
	FromReferralCode string `json:"from_referral_code,omitempty"`
}

type CreateStaticWalletData struct {
	WalletUUID string `json:"wallet_uuid"`
	UUID       string `json:"uuid"`
	Address    string `json:"address"`
	Network    string `json:"network"`
	Currency   string `json:"currency"`
	URL        string `json:"url"`
}

type CreateStaticWalletResponse struct {
	*HTTPResponse
	Result CreateStaticWalletData `json:"result"`
}

// CreateStaticWallet creates a static wallet.
//
// Details: https://doc.cryptomus.com/business/payments/creating-static
//
// Example:
//
//	result, err := sdk.CreateStaticWallet(&CreateStaticWalletRequest{
//		Currency: "USDT",
//		Network:  "tron",
//		OrderID:  "9d7f7b9b-3b7b-4b7b-9b7b-7b9d7f7b9b7b",
//		URLCallback: "https://example.com/callback",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) CreateStaticWallet(payload *CreateStaticWalletRequest) (*CreateStaticWalletResponse, error) {
	return sdk.CreateStaticWalletWithContext(context.Background(), payload)
}

// CreateStaticWalletWithContext creates a static wallet.
//
// Details: https://doc.cryptomus.com/business/payments/creating-static
//
// Example:
//
//	result, err := sdk.CreateStaticWalletWithContext(ctx, &CreateStaticWalletRequest{
//		Currency: "USDT",
//		Network:  "tron",
//		OrderID:  "9d7f7b9b-3b7b-4b7b-9b7b-7b9d7f7b9b7b",
//		URLCallback: "https://example.com/callback",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) CreateStaticWalletWithContext(ctx context.Context, payload *CreateStaticWalletRequest) (*CreateStaticWalletResponse, error) {
	var result CreateStaticWalletResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, payload)).
		SetBody(payload).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(CreateStaticWalletEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type GenerateQRCodeWalletRequest struct {
	WalletAddressUUID string `json:"wallet_address_uuid"`
}

type GenerateQRCodeWalletData struct {
	Image string `json:"image"`
}

type GenerateQRCodeWalletResponse struct {
	*HTTPResponse
	Result GenerateQRCodeWalletData `json:"result"`
}

// GenerateQRStaticWallet generates a QR code for a static wallet.
//
// Details: https://doc.cryptomus.com/business/payments/qr-code-pay-form
//
// You can generate a QR-code for a static wallet address or for an invoice address. Scanning it, the user will receive the address for depositing funds.
//
// Example:
//
//	result, err := sdk.GenerateQRStaticWallet(&GenerateQRCodeWalletRequest{
//		WalletAddressUUID: "9d7f7b9b-3b7b-4b7b-9b7b-7b9d7f7b9b7b",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) GenerateQRStaticWallet(payload *GenerateQRCodeWalletRequest) (*GenerateQRCodeWalletResponse, error) {
	return sdk.GenerateQRStaticWalletWithContext(context.Background(), payload)
}

// GenerateQRStaticWalletWithContext generates a QR code for a static wallet.
//
// Details: https://doc.cryptomus.com/business/payments/qr-code-pay-form
//
// You can generate a QR-code for a static wallet address or for an invoice address. Scanning it, the user will receive the address for depositing funds.
//
// Example:
//
//	result, err := sdk.GenerateQRStaticWalletWithContext(ctx, &GenerateQRCodeWalletRequest{
//		WalletAddressUUID: "9d7f7b9b-3b7b-4b7b-9b7b-7b9d7f7b9b7b",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) GenerateQRStaticWalletWithContext(ctx context.Context, payload *GenerateQRCodeWalletRequest) (*GenerateQRCodeWalletResponse, error) {
	var result GenerateQRCodeWalletResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, payload)).
		SetBody(payload).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(GenerateQRCodeWalletEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type GenerateQRCodeInvoiceRequest struct {
	MerchantPaymentUUID string `json:"merchant_payment_uuid"`
}

type GenerateQRCodeInvoiceData struct {
	Image string `json:"image"`
}

type GenerateQRCodeInvoiceResponse struct {
	*HTTPResponse
	Result GenerateQRCodeInvoiceData `json:"result"`
}

// GenerateQRCodeInvoice generates a QR code for an invoice.
//
// Details: https://doc.cryptomus.com/business/payments/qr-code-pay-form
//
// You can generate a QR-code for a static wallet address or for an invoice address. Scanning it, the user will receive the address for depositing funds.
//
// Example:
//
//	result, err := sdk.GenerateQRCodeInvoice(&GenerateQRCodeInvoiceRequest{
//		MerchantPaymentUUID: "9d7f7b9b-3b7b-4b7b-9b7b-7b9d7f7b9b7b",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) GenerateQRCodeInvoice(payload *GenerateQRCodeInvoiceRequest) (*GenerateQRCodeInvoiceResponse, error) {
	return sdk.GenerateQRCodeInvoiceWithContext(context.Background(), payload)
}

// GenerateQRCodeInvoiceWithContext generates a QR code for an invoice.
//
// Details: https://doc.cryptomus.com/business/payments/qr-code-pay-form
//
// You can generate a QR-code for a static wallet address or for an invoice address. Scanning it, the user will receive the address for depositing funds.
//
// Example:
//
//	result, err := sdk.GenerateQRCodeInvoiceWithContext(ctx, &GenerateQRCodeInvoiceRequest{
//		MerchantPaymentUUID: "9d7f7b9b-3b7b-4b7b-9b7b-7b9d7f7b9b7b",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) GenerateQRCodeInvoiceWithContext(ctx context.Context, payload *GenerateQRCodeInvoiceRequest) (*GenerateQRCodeInvoiceResponse, error) {
	var result GenerateQRCodeInvoiceResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, payload)).
		SetBody(payload).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(GenerateQRCodeInvoiceEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type BlockStaticWalletRequest struct {
	UUID          string `json:"uuid,omitempty"`
	OrderID       string `json:"order_id,omitempty"`
	IsForceRefund bool   `json:"is_force_refund"`
}

type BlockStaticWalletStatus string

const (
	BlockStaticWalletStatusBlocked  BlockStaticWalletStatus = "blocked"
	BlockStaticWalletStatusActive   BlockStaticWalletStatus = "active"
	BlockStaticWalletStatusInActive BlockStaticWalletStatus = "in_active"
)

type BlockStaticWalletData struct {
	UUID   string                  `json:"uuid"`
	Status BlockStaticWalletStatus `json:"status"`
}

type BlockStaticWalletResponse struct {
	*HTTPResponse
	Result BlockStaticWalletData `json:"result"`
}

// BlockStaticWallet blocks a static wallet.
//
// Details: https://doc.cryptomus.com/business/payments/block-wallet
//
// When you need to block your clients static wallet, all the further payments will not be credited to his balance. You can make a refund of this funds only once. The funds will be returned to the addresses from which they came.
//
// You need to pass one of the required parameters, if you pass both, the account will be identified by order_id
//
// Example:
//
//	result, err := sdk.BlockStaticWallet(&BlockStaticWalletRequest{
//		UUID: "9d7f7b9b-3b7b-4b7b-9b7b-7b9d7f7b9b7b",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) BlockStaticWallet(payload *BlockStaticWalletRequest) (*BlockStaticWalletResponse, error) {
	return sdk.BlockStaticWalletWithContext(context.Background(), payload)
}

// BlockStaticWalletWithContext blocks a static wallet.
//
// Details: https://doc.cryptomus.com/business/payments/block-wallet
//
// When you need to block your clients static wallet, all the further payments will not be credited to his balance. You can make a refund of this funds only once. The funds will be returned to the addresses from which they came.
//
// You need to pass one of the required parameters, if you pass both, the account will be identified by order_id
//
// Example:
//
//	result, err := sdk.BlockStaticWalletWithContext(ctx, &BlockStaticWalletRequest{
//		UUID: "9d7f7b9b-3b7b-4b7b-9b7b-7b9d7f7b9b7b",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) BlockStaticWalletWithContext(ctx context.Context, payload *BlockStaticWalletRequest) (*BlockStaticWalletResponse, error) {
	var result BlockStaticWalletResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, payload)).
		SetBody(payload).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(BlockStaticWalletEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type RefundPaymentOnBlockedAddressRequest struct {
	UUID    string `json:"uuid,omitempty"`
	OrderID string `json:"order_id,omitempty"`
	Address string `json:"address"`
}

type RefundPaymentOnBlockedAddressData struct {
	Commission string `json:"commission"`
	Amount     string `json:"amount"`
}

type RefundPaymentOnBlockedAddressResponse struct {
	*HTTPResponse
	Result RefundPaymentOnBlockedAddressData `json:"result"`
}

// RefundPaymentOnBlockedAddress refunds a payment on a blocked address.
//
// Details: https://doc.cryptomus.com/business/payments/refundblocked
//
// You can make a refund only once.
//
// To refund payments you need to pass either uuid or order_id, if you pass both, the static wallet will be identified by uuid
//
// Example:
//
//	result, err := sdk.RefundPaymentOnBlockedAddress(&RefundPaymentOnBlockedAddressRequest{
//		OrderID: "9d7f7b9b-3b7b-4b7b-9b7b-7b9d7f7b9b7b",
//		Address: "0x...",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) RefundPaymentOnBlockedAddress(payload *RefundPaymentOnBlockedAddressRequest) (*RefundPaymentOnBlockedAddressResponse, error) {
	return sdk.RefundPaymentOnBlockedAddressWithContext(context.Background(), payload)
}

// RefundPaymentOnBlockedAddressWithContext refunds a payment on a blocked address.
//
// Details: https://doc.cryptomus.com/business/payments/refundblocked
//
// You can make a refund only once.
//
// To refund payments you need to pass either uuid or order_id, if you pass both, the static wallet will be identified by uuid
//
// Example:
//
//	result, err := sdk.RefundPaymentOnBlockedAddressWithContext(ctx, &RefundPaymentOnBlockedAddressRequest{
//		OrderID: "9d7f7b9b-3b7b-4b7b-9b7b-7b9d7f7b9b7b",
//		Address: "0x...",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) RefundPaymentOnBlockedAddressWithContext(ctx context.Context, payload *RefundPaymentOnBlockedAddressRequest) (*RefundPaymentOnBlockedAddressResponse, error) {
	var result RefundPaymentOnBlockedAddressResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, payload)).
		SetBody(payload).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(RefundPaymentOnBlockedAddressEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type PaymentInformationRequest struct {
	UUID    string `json:"uuid,omitempty"`
	OrderID string `json:"order_id,omitempty"`
}

type PaymentInformationData struct {
	*CreateInvoiceData
}

type PaymentInformationResponse struct {
	*HTTPResponse
	Result PaymentInformationData `json:"result"`
}

// PaymentInformation returns payment information.
//
// Details: https://doc.cryptomus.com/business/payments/payment-information
//
// To get the invoice status you need to pass one of the required parameters, if you pass both, the account will be identified by order_id
//
// Example:
//
//	result, err := sdk.PaymentInformation(&PaymentInformationRequest{
//		UUID: "4b1b3b7b-7b1b-4b1b-7b1b-4b1b3b7b1b4b",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) PaymentInformation(payload *PaymentInformationRequest) (*PaymentInformationResponse, error) {
	return sdk.PaymentInformationWithContext(context.Background(), payload)
}

// PaymentInformationWithContext returns payment information.
//
// Details: https://doc.cryptomus.com/business/payments/payment-information
//
// To get the invoice status you need to pass one of the required parameters, if you pass both, the account will be identified by order_id
//
// Example:
//
//	result, err := sdk.PaymentInformationWithContext(ctx, &PaymentInformationRequest{
//		UUID: "4b1b3b7b-7b1b-4b1b-7b1b-4b1b3b7b1b4b",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) PaymentInformationWithContext(ctx context.Context, payload *PaymentInformationRequest) (*PaymentInformationResponse, error) {
	var result PaymentInformationResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, payload)).
		SetBody(payload).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(PaymentInformationEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type RefundRequest struct {
	Address    string `json:"address"`
	IsSubtract bool   `json:"is_subtract"`
	UUID       string `json:"uuid,omitempty"`
	OrderID    string `json:"order_id,omitempty"`
}

type RefundData struct{}

type RefundResponse struct {
	*HTTPResponse
	Result []RefundData `json:"result"`
}

// Refund refunds a payment.
//
// Details: https://doc.cryptomus.com/business/payments/refund
//
// Invoice is identified by order_id or uuid, if you pass both, the account will be identified by uuid
//
// Example:
//
//	result, err := sdk.Refund(&RefundRequest{
//		OrderID: "9d7f7b9b-3b7b-4b7b-9b7b-7b9d7f7b9b7b",
//		Address: "0x...",
//		IsSubtract: true,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) Refund(payload *RefundRequest) (*RefundResponse, error) {
	return sdk.RefundWithContext(context.Background(), payload)
}

// RefundWithContext refunds a payment.
//
// Details: https://doc.cryptomus.com/business/payments/refund
//
// Invoice is identified by order_id or uuid, if you pass both, the account will be identified by uuid
//
// Example:
//
//	result, err := sdk.RefundWithContext(ctx, &RefundRequest{
//		OrderID: "9d7f7b9b-3b7b-4b7b-9b7b-7b9d7f7b9b7b",
//		Address: "0x...",
//		IsSubtract: true,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) RefundWithContext(ctx context.Context, payload *RefundRequest) (*RefundResponse, error) {
	var result RefundResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, payload)).
		SetBody(payload).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(RefundEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type ResendWebhookRequest struct {
	UUID    string `json:"uuid,omitempty"`
	OrderID string `json:"order_id,omitempty"`
}

type ResendWebhookData struct{}

type ResendWebhookResponse struct {
	*HTTPResponse
	Result []ResendWebhookData `json:"result"`
}

// ResendWebhook resends a webhook.
//
// Details: https://doc.cryptomus.com/business/payments/resend-webhook
//
// Resend the webhook by invoice. You can resend the webhook only for finalized invoices, that is, invoices in statuses: wrong_amount, paid, paid_over.
//
// To resend the webhook on the invoice, the url_callback must be specified at the time of invoice creation.
//
// Example:
//
//	result, err := sdk.ResendWebhook(&ResendWebhookRequest{
//		OrderID: "9d7f7b9b-3b7b-4b7b-9b7b-7b9d7f7b9b7b",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) ResendWebhook(payload *ResendWebhookRequest) (*ResendWebhookResponse, error) {
	return sdk.ResendWebhookWithContext(context.Background(), payload)
}

// ResendWebhookWithContext resends a webhook.
//
// Details: https://doc.cryptomus.com/business/payments/resend-webhook
//
// Resend the webhook by invoice. You can resend the webhook only for finalized invoices, that is, invoices in statuses: wrong_amount, paid, paid_over.
//
// To resend the webhook on the invoice, the url_callback must be specified at the time of invoice creation.
//
// Example:
//
//	result, err := sdk.ResendWebhookWithContext(ctx, &ResendWebhookRequest{
//		OrderID: "9d7f7b9b-3b7b-4b7b-9b7b-7b9d7f7b9b7b",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) ResendWebhookWithContext(ctx context.Context, payload *ResendWebhookRequest) (*ResendWebhookResponse, error) {
	var result ResendWebhookResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, payload)).
		SetBody(payload).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(ResendWebhookEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type PaymentStatus string

const (
	PaymentStatusProcess            PaymentStatus = "process"              // Payment in processing.
	PaymentStatusCheck              PaymentStatus = "check"                // Waiting for the transaction to appear on the blockchain.
	PaymentStatusConfirmCheck       PaymentStatus = "confirm_check"        // We have seen the transaction in the blockchain and are waiting for the required number of network confirmations.
	PaymentStatusWrongAmountWaiting PaymentStatus = "wrong_amount_waiting" // The client paid less than required, with the possibility of an additional payment.
	PaymentStatusPaid               PaymentStatus = "paid"                 // The payment was successful and the client paid exactly as much as required.
	PaymentStatusPaidOver           PaymentStatus = "paid_over"            // The payment was successful and client paid more than required.
	PaymentStatusFail               PaymentStatus = "fail"                 // Payment error.
	PaymentStatusWrongAmount        PaymentStatus = "wrong_amount"         // The client paid less than required.
	PaymentStatusCancel             PaymentStatus = "cancel"               // Payment cancelled, the client did not pay.
	PaymentStatusSystemFail         PaymentStatus = "system_fail"          // A system error has occurred.
	PaymentStatusRefundProcess      PaymentStatus = "refund_process"       // The refund is being processed.
	PaymentStatusRefundFail         PaymentStatus = "refund_fail"          // An error occurred during the refund.
	PaymentStatusRefundPaid         PaymentStatus = "refund_paid"          // The refund was successful.
	PaymentStatusLocked             PaymentStatus = "locked"               // Funds are locked due to the AML program.
)

type TestingWebhookRequest struct {
	URLCallback string        `json:"url_callback"`
	Currency    string        `json:"currency"`
	Network     string        `json:"network"`
	UUID        string        `json:"uuid,omitempty"`
	OrderID     string        `json:"order_id,omitempty"`
	Status      PaymentStatus `json:"status"`
}

type TestingWebhookData struct{}

type TestingWebhookResponse struct {
	*HTTPResponse
	Result []TestingWebhookData `json:"result"`
}

// TestingWebhook tests a webhook.
//
// Details: https://doc.cryptomus.com/business/payments/testing-webhook
//
// To validate the signature from the webhook data array, use the payment API key.
//
// To ensure that you are correctly receiving webhooks and can validate the signature, you should use this method to test webhooks for payment.
// Please note that no data is saved to the database, and any data received in the webhook is only stored in an array for testing purposes to ensure the correctness of the signature and to test the retrieval of this array from us.
//
// To test a webhook with an existing invoice, please provide its uuid or order ID. If these parameters are not provided, the webhook will be sent with a test invoice.
//
// You may to pass one of the uuid or order_id parameters, if you pass both, the account will be identified by uuid
//
// Example:
//
//	result, err := sdk.TestingWebhook(&TestingWebhookRequest{
//		UUID:        "4b1b3b7b-7b1b-4b1b-7b1b-4b1b3b7b1b4b",
//		URLCallback: "https://example.com/callback",
//		Currency:    "USD",
//		Network:     "btc",
//		Status:      PaymentStatusPaid,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) TestingWebhook(payload *TestingWebhookRequest) (*TestingWebhookResponse, error) {
	return sdk.TestingWebhookWithContext(context.Background(), payload)
}

// TestingWebhookWithContext tests a webhook.
//
// Details: https://doc.cryptomus.com/business/payments/testing-webhook
//
// To validate the signature from the webhook data array, use the payment API key.
//
// To ensure that you are correctly receiving webhooks and can validate the signature, you should use this method to test webhooks for payment.
// Please note that no data is saved to the database, and any data received in the webhook is only stored in an array for testing purposes to ensure the correctness of the signature and to test the retrieval of this array from us.
//
// To test a webhook with an existing invoice, please provide its uuid or order ID. If these parameters are not provided, the webhook will be sent with a test invoice.
//
// You may to pass one of the uuid or order_id parameters, if you pass both, the account will be identified by uuid
//
// Example:
//
//	result, err := sdk.TestingWebhookWithContext(ctx, &TestingWebhookRequest{
//		UUID:        "4b1b3b7b-7b1b-4b1b-7b1b-4b1b3b7b1b4b",
//		URLCallback: "https://example.com/callback",
//		Currency:    "USD",
//		Network:     "btc",
//		Status:      PaymentStatusPaid,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) TestingWebhookWithContext(ctx context.Context, payload *TestingWebhookRequest) (*TestingWebhookResponse, error) {
	var result TestingWebhookResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, payload)).
		SetBody(payload).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(TestingWebhookPaymentEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type PaymentListOfServiceLimit struct {
	MinAmount string `json:"min_amount"`
	MaxAmount string `json:"max_amount"`
}

type PaymentListOfServiceCommission struct {
	FeeAmount string `json:"fee_amount"`
	Percent   string `json:"percent"`
}

type PaymentListOfServicesData struct {
	Network     string                         `json:"network"`
	Currency    string                         `json:"currency"`
	IsAvailable bool                           `json:"is_available"`
	Limit       PaymentListOfServiceLimit      `json:"limit"`
	Commission  PaymentListOfServiceCommission `json:"commission"`
}

type PaymentListOfServicesResponse struct {
	*HTTPResponse
	Result []PaymentListOfServicesData `json:"result"`
}

// PaymentListOfServices returns a list of services.
//
// Details: https://doc.cryptomus.com/business/payments/list-of-services
//
// Returns a list of available payment services. Payment services store settings that are taken into account when creating an invoice. For example. currencies, networks, minimum and maximum limits, commissions.
//
// Example:
//
//	result, err := sdk.PaymentListOfServices()
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) PaymentListOfServices() (*PaymentListOfServicesResponse, error) {
	return sdk.PaymentListOfServicesWithContext(context.Background())
}

// PaymentListOfServicesWithContext returns a list of services.
//
// Details: https://doc.cryptomus.com/business/payments/list-of-services
//
// Returns a list of available payment services. Payment services store settings that are taken into account when creating an invoice. For example. currencies, networks, minimum and maximum limits, commissions.
//
// Example:
//
//	result, err := sdk.PaymentListOfServicesWithContext(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) PaymentListOfServicesWithContext(ctx context.Context) (*PaymentListOfServicesResponse, error) {
	var result PaymentListOfServicesResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, "")).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(PaymentListOfServicesEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type PaymentHistoryRequest struct {
	DateFrom string `json:"date_from,omitempty"`
	DateTo   string `json:"date_to,omitempty"`
	Cursor   string `json:"-"`
}

type PaymentHistoryData struct {
	Items    []CreateInvoiceData `json:"items"`
	Paginate *Pagination         `json:"paginate"`
}

type PaymentHistoryResponse struct {
	*HTTPResponse
	Result PaymentHistoryData `json:"result"`
}

// PaymentHistory returns payment history.
//
// Details: https://doc.cryptomus.com/business/payments/payment-history
//
// To get next/previous page entries, specify the next/previous cursor hash
//
// Example:
//
//	result, err := sdk.PaymentHistory(&PaymentHistoryRequest{
//		DateFrom: "2021-01-01 00:00:00",
//		DateTo:   "2021-12-31 23:59:59",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) PaymentHistory(payload *PaymentHistoryRequest) (*PaymentHistoryResponse, error) {
	return sdk.PaymentHistoryWithContext(context.Background(), payload)
}

// PaymentHistoryWithContext returns payment history.
//
// Details: https://doc.cryptomus.com/business/payments/payment-history
//
// To get next/previous page entries, specify the next/previous cursor hash
//
// Example:
//
//	result, err := sdk.PaymentHistoryWithContext(ctx, &PaymentHistoryRequest{
//		DateFrom: "2021-01-01 00:00:00",
//		DateTo:   "2021-12-31 23:59:59",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) PaymentHistoryWithContext(ctx context.Context, payload *PaymentHistoryRequest) (*PaymentHistoryResponse, error) {
	var result PaymentHistoryResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, payload)).
		SetBody(payload).
		SetQueryParam("cursor", payload.Cursor).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(PaymentHistoryEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}
