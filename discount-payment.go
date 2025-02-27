package cryptomus

import (
	"context"
)

type Discount struct {
	Currency string `json:"currency"`
	Network  string `json:"network"`
	Discount int    `json:"discount"`
}

type ListOfDiscount []Discount

type ListOfDiscountResponse struct {
	*HTTPResponse
	Result ListOfDiscount `json:"result"`
}

func (sdk *Cryptomus) ListOfDiscount() (*ListOfDiscountResponse, error) {
	return sdk.ListOfDiscountWithContext(context.Background())
}

func (sdk *Cryptomus) ListOfDiscountWithContext(ctx context.Context) (*ListOfDiscountResponse, error) {

	var result ListOfDiscountResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, "")).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(ListOfDiscountsEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type SetDiscountToPaymentMethodRequest struct {
	Currency string `json:"currency"`
	Network  string `json:"network"`
	Discount int    `json:"discount_percent"`
}

type SetDiscountToPaymentMethodResponse struct {
	*HTTPResponse
	Result Discount `json:"result,omitempty"`
}

func (sdk *Cryptomus) SetDiscountToPaymentMethod(payload SetDiscountToPaymentMethodRequest) (*SetDiscountToPaymentMethodResponse, error) {
	return sdk.SetDiscountToPaymentMethodWithContext(context.Background(), payload)
}

func (sdk *Cryptomus) SetDiscountToPaymentMethodWithContext(ctx context.Context, payload SetDiscountToPaymentMethodRequest) (*SetDiscountToPaymentMethodResponse, error) {

	var result SetDiscountToPaymentMethodResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, payload)).
		SetBody(payload).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(SetDiscountToPaymentMethodEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}
