package cryptomus_sdk_go

import (
	"context"
	"time"
)

type RecurringPaymentPeriod string

const (
	RecurringPaymentPeriodWeekly     RecurringPaymentPeriod = "weekly"
	RecurringPaymentPeriodMonthly    RecurringPaymentPeriod = "monthly"
	RecurringPaymentPeriodThreeMonth RecurringPaymentPeriod = "three_month"
)

type CreateRecurringPaymentRequest struct {
	Amount         string                 `json:"amount"`
	Currency       string                 `json:"currency"`
	Name           string                 `json:"name"`
	Period         RecurringPaymentPeriod `json:"period"`
	ToCurrency     string                 `json:"to_currency,omitempty"`
	OrderID        string                 `json:"order_id,omitempty"`
	URLCallback    string                 `json:"url_callback,omitempty"`
	DiscountDays   int                    `json:"discount_days,omitempty"`
	DiscountAmount string                 `json:"discount_amount,omitempty"`
	AdditionalData string                 `json:"additional_data,omitempty"`
}

type CreateRecurringPaymentResponse struct {
	*HTTPResponse
	Result RecurringPaymentData `json:"result,omitempty"`
}

type RecurringPaymentData struct {
	UUID           string     `json:"uuid"`
	Name           string     `json:"name"`
	OrderID        string     `json:"order_id,omitempty"`
	Amount         string     `json:"amount"`
	Currency       string     `json:"currency"`
	PayerCurrency  string     `json:"payer_currency,omitempty"`
	PayerAmountUSD string     `json:"payer_amount_usd"`
	PayerAmount    string     `json:"payer_amount,omitempty"`
	URLCallback    string     `json:"url_callback,omitempty"`
	Period         string     `json:"period"`
	Status         string     `json:"status"`
	URL            string     `json:"url"`
	LastPayOff     *time.Time `json:"last_pay_off,omitempty"`
}

func (sdk *Cryptomus) CreateRecurringPayment(ctx context.Context, request CreateRecurringPaymentRequest) (*CreateRecurringPaymentResponse, error) {
	var result CreateRecurringPaymentResponse

	reqByte, err := ToJSON(request)
	if err != nil {
		return nil, err
	}

	req := sdk.HttpClient.NewRequest().
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, string(reqByte))).
		SetHeader("Content-Type", "application/json").
		SetBody(reqByte).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(CreateRecurringPaymentEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type RecurringPaymentInformationRequest struct {
	UUID    string `json:"uuid,omitempty"`
	OrderID string `json:"order_id,omitempty"`
}

type RecurringPaymentInformationResponse struct {
	*HTTPResponse
	Result RecurringPaymentData `json:"result,omitempty"`
}

func (sdk *Cryptomus) RecurringPaymentInformation(ctx context.Context, request RecurringPaymentInformationRequest) (*RecurringPaymentInformationResponse, error) {
	var result RecurringPaymentInformationResponse

	reqByte, err := ToJSON(request)
	if err != nil {
		return nil, err
	}

	req := sdk.HttpClient.NewRequest().
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, string(reqByte))).
		SetHeader("Content-Type", "application/json").
		SetBody(reqByte).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(RecurringPaymentInformationEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type ListRecurringPaymentsRequest struct {
	Cursor string `json:"cursor,omitempty"`
}

type ListRecurringPaymentsResponse struct {
	*HTTPResponse
	Result ListRecurringPaymentsData `json:"result,omitempty"`
}

type ListRecurringPaymentsData struct {
	Items    []RecurringPaymentData `json:"items"`
	Paginate *Pagination            `json:"paginate"`
}

func (sdk *Cryptomus) ListRecurringPayments(ctx context.Context, request ListRecurringPaymentsRequest) (*ListRecurringPaymentsResponse, error) {
	var result ListRecurringPaymentsResponse

	reqByte, err := ToJSON(request)
	if err != nil {
		return nil, err
	}

	req := sdk.HttpClient.NewRequest().
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, string(reqByte))).
		SetHeader("Content-Type", "application/json").
		SetQueryParam("cursor", request.Cursor).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(ListRecurringPaymentsEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type CancelRecurringPaymentRequest struct {
	UUID    string `json:"uuid,omitempty"`
	OrderID string `json:"order_id,omitempty"`
}

type CancelRecurringPaymentResponse struct {
	*HTTPResponse
	Result RecurringPaymentData `json:"result,omitempty"`
}

func (sdk *Cryptomus) CancelRecurringPayment(ctx context.Context, request CancelRecurringPaymentRequest) (*CancelRecurringPaymentResponse, error) {
	var result CancelRecurringPaymentResponse

	reqByte, err := ToJSON(request)
	if err != nil {
		return nil, err
	}

	req := sdk.HttpClient.NewRequest().
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, string(reqByte))).
		SetHeader("Content-Type", "application/json").
		SetBody(reqByte).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(CancelRecurringPaymentEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}
