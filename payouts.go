package cryptomus

import (
	"context"
	"time"
)

type PayoutStatus string

const (
	PayoutStatusProcess    PayoutStatus = "process"
	PayoutStatusCheck      PayoutStatus = "check"
	PayoutStatusPaid       PayoutStatus = "paid"
	PayoutStatusFail       PayoutStatus = "fail"
	PayoutStatusCancel     PayoutStatus = "cancel"
	PayoutStatusSystemFail PayoutStatus = "system_fail"
)

type CreatePayoutRequest struct {
	Amount       string `json:"amount"`
	Currency     string `json:"currency" validate:"required"`
	OrderID      string `json:"order_id"`
	Address      string `json:"address"`
	IsSubtract   bool   `json:"is_subtract"`
	Network      string `json:"network"`
	URLCallback  string `json:"url_callback"`
	ToCurrency   string `json:"to_currency"`
	CourseSource string `json:"course_source"`
	FromCurrency string `json:"from_currency"`
	Priority     string `json:"priority"`
	Memo         string `json:"memo"`
}

type PayoutData struct {
	UUID          string       `json:"uuid"`
	Amount        string       `json:"amount"`
	Currency      string       `json:"currency"`
	Network       string       `json:"network"`
	Address       string       `json:"address"`
	TxID          string       `json:"txid"`
	Status        PayoutStatus `json:"status"`
	IsFinal       bool         `json:"is_final"`
	Balance       int64        `json:"balance"`
	PayerCurrency string       `json:"payer_currency"`
	PayerAmount   int64        `json:"payer_amount"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}

type CreatePayoutResponse struct {
	*HTTPResponse
	Result *PayoutData `json:"result"`
}

func (sdk *Cryptomus) CreatePayout(payload *CreatePayoutRequest) (*CreatePayoutResponse, error) {
	return sdk.CreatePayoutWithContext(context.Background(), payload)
}

func (sdk *Cryptomus) CreatePayoutWithContext(ctx context.Context, payload *CreatePayoutRequest) (*CreatePayoutResponse, error) {
	var result CreatePayoutResponse

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

	if _, err := req.Post(CreatePayoutEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type PayoutInformationRequest struct {
	UUID    string `json:"uuid,omitempty"`
	OrderID string `json:"order_id,omitempty"`
}

type PayoutInformationResponse struct {
	*HTTPResponse
	Result *PayoutData `json:"result"`
}

func (sdk *Cryptomus) PayoutInformation(payload *PayoutInformationRequest) (*PayoutInformationResponse, error) {
	return sdk.PayoutInformationWithContext(context.Background(), payload)
}

func (sdk *Cryptomus) PayoutInformationWithContext(ctx context.Context, payload *PayoutInformationRequest) (*PayoutInformationResponse, error) {
	var result PayoutInformationResponse

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

	if _, err := req.Post(PayoutInformationEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type PayoutHistoryRequest struct {
	DateFrom string `json:"date_from,omitempty"`
	DateTo   string `json:"date_to,omitempty"`
}

type PayoutHistoryData struct {
	MerchantUUID string        `json:"merchant_uuid"`
	Items        []*PayoutData `json:"items"`
	Paginate     *Pagination   `json:"paginate"`
}

type PayoutHistoryResponse struct {
	*HTTPResponse
	Result *PayoutHistoryData `json:"result"`
}

func (sdk *Cryptomus) PayoutHistory(payload *PayoutHistoryRequest) (*PayoutHistoryResponse, error) {
	return sdk.PayoutHistoryWithContext(context.Background(), payload)
}

func (sdk *Cryptomus) PayoutHistoryWithContext(ctx context.Context, payload *PayoutHistoryRequest) (*PayoutHistoryResponse, error) {
	var result PayoutHistoryResponse

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

	if _, err := req.Post(PayoutHistoryEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type PayoutListOfServiceLimit struct {
	MinAmount string `json:"min_amount"`
	MaxAmount string `json:"max_amount"`
}

type PayoutListOfServiceCommission struct {
	FeeAmount string `json:"fee_amount"`
	Percent   string `json:"percent"`
}

type PayoutListOfServicesData struct {
	Network     string                         `json:"network"`
	Currency    string                         `json:"currency"`
	IsAvailable bool                           `json:"is_available"`
	Limit       PaymentListOfServiceLimit      `json:"limit"`
	Commission  PaymentListOfServiceCommission `json:"commission"`
}

type PayoutListOfServicesResponse struct {
	*HTTPResponse
	Result []*PayoutListOfServicesData `json:"result"`
}

func (sdk *Cryptomus) PayoutListOfServices() (*PayoutListOfServicesResponse, error) {
	return sdk.PayoutListOfServicesWithContext(context.Background())
}

func (sdk *Cryptomus) PayoutListOfServicesWithContext(ctx context.Context) (*PayoutListOfServicesResponse, error) {
	var result PayoutListOfServicesResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, "")).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Get(PayoutListOfServicesEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type TransferToPersonalWalletRequest struct {
	Amount   string `json:"amount" validate:"required"`
	Currency string `json:"currency" validate:"required"`
}

type TransferToPersonalWalletData struct {
	UserWalletTransactionUUID string `json:"user_wallet_transaction_uuid"`
	UserWalletBalance         string `json:"user_wallet_balance"`
	MerchantTransactionUUID   string `json:"merchant_transaction_uuid"`
	MerchantBalance           string `json:"merchant_balance"`
}

type TransferToPersonalWalletResponse struct {
	*HTTPResponse
	Result *TransferToPersonalWalletData `json:"result"`
}

func (sdk *Cryptomus) TransferToPersonalWallet(payload *TransferToPersonalWalletRequest) (*TransferToPersonalWalletResponse, error) {
	return sdk.TransferToPersonalWalletWithContext(context.Background(), payload)
}

func (sdk *Cryptomus) TransferToPersonalWalletWithContext(ctx context.Context, payload *TransferToPersonalWalletRequest) (*TransferToPersonalWalletResponse, error) {
	var result TransferToPersonalWalletResponse

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

	if _, err := req.Post(TransferToPersonalWalletEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type TransferToBusinessWalletRequest struct {
	Amount   string `json:"amount" validate:"required"`
	Currency string `json:"currency" validate:"required"`
}

type TransferToBusinessWalletData struct {
	UserWalletTransactionUUID string `json:"user_wallet_transaction_uuid"`
	UserWalletBalance         string `json:"user_wallet_balance"`
	MerchantTransactionUUID   string `json:"merchant_transaction_uuid"`
	MerchantBalance           string `json:"merchant_balance"`
}

type TransferToBusinessWalletResponse struct {
	*HTTPResponse
	Result *TransferToBusinessWalletData `json:"result"`
}

func (sdk *Cryptomus) TransferToBusinessWallet(payload *TransferToBusinessWalletRequest) (*TransferToBusinessWalletResponse, error) {
	return sdk.TransferToBusinessWalletWithContext(context.Background(), payload)
}

func (sdk *Cryptomus) TransferToBusinessWalletWithContext(ctx context.Context, payload *TransferToBusinessWalletRequest) (*TransferToBusinessWalletResponse, error) {
	var result TransferToBusinessWalletResponse

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

	if _, err := req.Post(TransferToBusinessWalletEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}
