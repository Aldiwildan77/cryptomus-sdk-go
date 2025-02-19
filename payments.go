package cryptomus_sdk_go

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

func (sdk *Cryptomus) CreateInvoice(ctx context.Context, payload *CreateInvoiceRequest) (*CreateInvoiceResponse, error) {
	var result CreateInvoiceResponse

	payloadByte, err := ToJSON(payload)
	if err != nil {
		return nil, err
	}

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, string(payloadByte))).
		SetHeader("Content-Type", "application/json").
		SetSuccessResult(&result).
		SetErrorResult(&result).
		SetBody(payloadByte)

	if _, err := req.Post(CreateInvoiceEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}
