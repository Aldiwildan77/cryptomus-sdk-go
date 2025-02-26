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

func (sdk *Cryptomus) ListOfDiscount(ctx context.Context) (*ListOfDiscountResponse, error) {

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

func (sdk *Cryptomus) SetDiscountToPaymentMethod(ctx context.Context, request SetDiscountToPaymentMethodRequest) (*SetDiscountToPaymentMethodResponse, error) {

	var result SetDiscountToPaymentMethodResponse

	reqByte, err := ToJSON(request)
	if err != nil {
		return nil, err
	}

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, string(reqByte))).
		SetBody(request).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(SetDiscountToPaymentMethodEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}
