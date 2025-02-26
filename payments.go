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
		SetSuccessResult(&result).
		SetErrorResult(&result).
		SetBody(payloadByte)

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

func (sdk *Cryptomus) CreateStaticWallet(ctx context.Context, payload *CreateStaticWalletRequest) (*CreateStaticWalletResponse, error) {
	var result CreateStaticWalletResponse

	payloadByte, err := ToJSON(payload)
	if err != nil {
		return nil, err
	}

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, string(payloadByte))).
		SetSuccessResult(&result).
		SetErrorResult(&result).
		SetBody(payloadByte)

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

func (sdk *Cryptomus) GenerateQRStaticWallet(ctx context.Context, payload *GenerateQRCodeWalletRequest) (*GenerateQRCodeWalletResponse, error) {
	var result GenerateQRCodeWalletResponse

	payloadByte, err := ToJSON(payload)
	if err != nil {
		return nil, err
	}

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, string(payloadByte))).
		SetSuccessResult(&result).
		SetErrorResult(&result).
		SetBody(payloadByte)

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

func (sdk *Cryptomus) GenerateQRCodeInvoice(ctx context.Context, payload *GenerateQRCodeInvoiceRequest) (*GenerateQRCodeInvoiceResponse, error) {
	var result GenerateQRCodeInvoiceResponse

	payloadByte, err := ToJSON(payload)
	if err != nil {
		return nil, err
	}

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, string(payloadByte))).
		SetSuccessResult(&result).
		SetErrorResult(&result).
		SetBody(payloadByte)

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

func (sdk *Cryptomus) BlockStaticWallet(ctx context.Context, payload *BlockStaticWalletRequest) (*BlockStaticWalletResponse, error) {
	var result BlockStaticWalletResponse

	payloadByte, err := ToJSON(payload)
	if err != nil {
		return nil, err
	}

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, string(payloadByte))).
		SetSuccessResult(&result).
		SetErrorResult(&result).
		SetBody(payloadByte)

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

func (sdk *Cryptomus) RefundPaymentOnBlockedAddress(ctx context.Context, payload *RefundPaymentOnBlockedAddressRequest) (*RefundPaymentOnBlockedAddressResponse, error) {
	var result RefundPaymentOnBlockedAddressResponse

	payloadByte, err := ToJSON(payload)
	if err != nil {
		return nil, err
	}

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, string(payloadByte))).
		SetSuccessResult(&result).
		SetErrorResult(&result).
		SetBody(payloadByte)

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

func (sdk *Cryptomus) PaymentInformation(ctx context.Context, payload *PaymentInformationRequest) (*PaymentInformationResponse, error) {
	var result PaymentInformationResponse

	payloadByte, err := ToJSON(payload)
	if err != nil {
		return nil, err
	}

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, string(payloadByte))).
		SetSuccessResult(&result).
		SetErrorResult(&result).
		SetBody(payloadByte)

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

func (sdk *Cryptomus) Refund(ctx context.Context, payload *RefundRequest) (*RefundResponse, error) {
	var result RefundResponse

	payloadByte, err := ToJSON(payload)
	if err != nil {
		return nil, err
	}

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, string(payloadByte))).
		SetSuccessResult(&result).
		SetErrorResult(&result).
		SetBody(payloadByte)

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

func (sdk *Cryptomus) ResendWebhook(ctx context.Context, payload *ResendWebhookRequest) (*ResendWebhookResponse, error) {
	var result ResendWebhookResponse

	payloadByte, err := ToJSON(payload)
	if err != nil {
		return nil, err
	}

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, string(payloadByte))).
		SetSuccessResult(&result).
		SetErrorResult(&result).
		SetBody(payloadByte)

	if _, err := req.Post(ResendWebhookEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type PaymentStatus string

const (
	PaymentStatusProcess       PaymentStatus = "process"
	PaymentStatusCheck         PaymentStatus = "check"
	PaymentStatusPaid          PaymentStatus = "paid"
	PaymentStatusPaidOver      PaymentStatus = "paid_over"
	PaymentStatusFail          PaymentStatus = "fail"
	PaymentStatusWrongAmount   PaymentStatus = "wrong_amount"
	PaymentStatusCancel        PaymentStatus = "cancel"
	PaymentStatusSystemFail    PaymentStatus = "system_fail"
	PaymentStatusRefundProcess PaymentStatus = "refund_process"
	PaymentStatusRefundFail    PaymentStatus = "refund_fail"
	PaymentStatusRefundPaid    PaymentStatus = "refund_paid"
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

func (sdk *Cryptomus) TestingWebhook(ctx context.Context, payload *TestingWebhookRequest) (*TestingWebhookResponse, error) {
	var result TestingWebhookResponse

	payloadByte, err := ToJSON(payload)
	if err != nil {
		return nil, err
	}

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, string(payloadByte))).
		SetSuccessResult(&result).
		SetErrorResult(&result).
		SetBody(payloadByte)

	if _, err := req.Post(TestingWebhookEndpoint.URL()); err != nil {
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

func (sdk *Cryptomus) PaymentListOfServices(ctx context.Context) (*PaymentListOfServicesResponse, error) {
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
}

type PaymentHistoryData struct {
	Items    []CreateInvoiceData `json:"items"`
	Paginate *Pagination         `json:"paginate"`
}

type PaymentHistoryResponse struct {
	*HTTPResponse
	Result PaymentHistoryData `json:"result"`
}

func (sdk *Cryptomus) PaymentHistory(ctx context.Context, payload *PaymentHistoryRequest) (*PaymentHistoryResponse, error) {
	var result PaymentHistoryResponse

	payloadByte, err := ToJSON(payload)
	if err != nil {
		return nil, err
	}

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, string(payloadByte))).
		SetSuccessResult(&result).
		SetErrorResult(&result).
		SetBody(payloadByte)

	if _, err := req.Post(PaymentHistoryEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}
