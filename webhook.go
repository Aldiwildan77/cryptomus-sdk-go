package cryptomus

import "context"

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

type TestingWebhookPaymentRequest struct {
	URLCallback string        `json:"url_callback"`
	Currency    string        `json:"currency"`
	Network     string        `json:"network"`
	UUID        string        `json:"uuid,omitempty"`
	OrderID     string        `json:"order_id,omitempty"`
	Status      PaymentStatus `json:"status"`
}

type TestingWebhookPaymentData struct{}

type TestingWebhookPaymentResponse struct {
	*HTTPResponse
	Result []TestingWebhookPaymentData `json:"result"`
}

// TestingWebhookPayment tests a webhook.
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
//	result, err := sdk.TestingWebhookPayment(&TestingWebhookRequest{
//		UUID:        "4b1b3b7b-7b1b-4b1b-7b1b-4b1b3b7b1b4b",
//		URLCallback: "https://example.com/callback",
//		Currency:    "USD",
//		Network:     "btc",
//		Status:      PaymentStatusPaid,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) TestingWebhookPayment(payload *TestingWebhookPaymentRequest) (*TestingWebhookPaymentResponse, error) {
	return sdk.TestingWebhookPaymentWithContext(context.Background(), payload)
}

// TestingWebhookPaymentWithContext tests a webhook.
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
//	result, err := sdk.TestingWebhookPaymentWithContext(ctx, &TestingWebhookRequest{
//		UUID:        "4b1b3b7b-7b1b-4b1b-7b1b-4b1b3b7b1b4b",
//		URLCallback: "https://example.com/callback",
//		Currency:    "USD",
//		Network:     "btc",
//		Status:      PaymentStatusPaid,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) TestingWebhookPaymentWithContext(ctx context.Context, payload *TestingWebhookPaymentRequest) (*TestingWebhookPaymentResponse, error) {
	var result TestingWebhookPaymentResponse

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

type TestingWebhookPayoutRequest struct {
	URLCallback string       `json:"url_callback"`
	Currency    string       `json:"currency"`
	Network     string       `json:"network"`
	UUID        string       `json:"uuid,omitempty"`
	OrderID     string       `json:"order_id,omitempty"`
	Status      PayoutStatus `json:"status"`
}

type TestingWebhookPayoutData struct{}

type TestingWebhookPayoutResponse struct {
	*HTTPResponse
	Result []TestingWebhookPayoutData `json:"result"`
}

// TestingWebhookPayout tests a webhook.
//
// Details: https://doc.cryptomus.com/business/payments/testing-webhook
//
// To validate the signature from the webhook data array, use the payment API key.
//
// You may to pass one of the uuid or order_id parameters, if you pass both, the account will be identified by uuid

// Example:
//
//	result, err := sdk.TestingWebhookPayout(&TestingWebhookPayoutRequest{
//		UUID:        "4b1b3b7b-7b1b-4b1b-7b1b-4b1b3b7b1b4b",
//		URLCallback: "https://example.com/callback",
//		Currency:    "USD",
//		Network:     "btc",
//		Status:      PayoutStatusPaid,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) TestingWebhookPayout(payload *TestingWebhookPayoutRequest) (*TestingWebhookPayoutResponse, error) {
	return sdk.TestingWebhookPayoutWithContext(context.Background(), payload)
}

// TestingWebhookPayoutWithContext tests a webhook.
//
// Details: https://doc.cryptomus.com/business/payments/testing-webhook
//
// To validate the signature from the webhook data array, use the payment API key.
//
// You may to pass one of the uuid or order_id parameters, if you pass both, the account will be identified by uuid

// Example:
//
//	result, err := sdk.TestingWebhookPayoutWithContext(ctx, &TestingWebhookPayoutRequest{
//		UUID:        "4b1b3b7b-7b1b-4b1b-7b1b-4b1b3b7b1b4b",
//		URLCallback: "https://example.com/callback",
//		Currency:    "USD",
//		Network:     "btc",
//		Status:      PayoutStatusPaid,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) TestingWebhookPayoutWithContext(ctx context.Context, payload *TestingWebhookPayoutRequest) (*TestingWebhookPayoutResponse, error) {
	var result TestingWebhookPayoutResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, payload)).
		SetBody(payload).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(TestingWebhookPayoutEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type TestingWebhookWalletRequest struct {
	URLCallback string        `json:"url_callback"`
	Currency    string        `json:"currency"`
	Network     string        `json:"network"`
	UUID        string        `json:"uuid,omitempty"`
	OrderID     string        `json:"order_id,omitempty"`
	Status      PaymentStatus `json:"status"`
}

type TestingWebhookWalletData struct{}

type TestingWebhookWalletResponse struct {
	*HTTPResponse
	Result []TestingWebhookWalletData `json:"result"`
}

// TestingWebhookWallet tests a webhook.
//
// Details: https://doc.cryptomus.com/business/payments/testing-webhook
//
// To validate the signature from the webhook data array, use the payment API key.
//
// Example:
//
//	result, err := sdk.TestingWebhookWallet(&TestingWebhookWalletRequest{
//		UUID:        "4b1b3b7b-7b1b-4b1b-7b1b-4b1b3b7b1b4b",
//		URLCallback: "https://example.com/callback",
//		Currency:    "USD",
//		Network:     "btc",
//		Status:      PaymentStatusPaid,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) TestingWebhookWallet(payload *TestingWebhookWalletRequest) (*TestingWebhookWalletResponse, error) {
	return sdk.TestingWebhookWalletWithContext(context.Background(), payload)
}

// TestingWebhookWalletWithContext tests a webhook.
//
// Details: https://doc.cryptomus.com/business/payments/testing-webhook
//
// To validate the signature from the webhook data array, use the payment API key.
//
// Example:
//
//	result, err := sdk.TestingWebhookWalletWithContext(ctx, &TestingWebhookWalletRequest{
//		UUID:        "4b1b3b7b-7b1b-4b1b-7b1b-4b1b3b7b1b4b",
//		URLCallback: "https://example.com/callback",
//		Currency:    "USD",
//		Network:     "btc",
//		Status:      PaymentStatusPaid,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) TestingWebhookWalletWithContext(ctx context.Context, payload *TestingWebhookWalletRequest) (*TestingWebhookWalletResponse, error) {
	var result TestingWebhookWalletResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, payload)).
		SetBody(payload).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(TestingWebhookWalletEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}
