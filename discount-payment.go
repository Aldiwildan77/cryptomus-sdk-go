package cryptomus

import (
	"context"
)

type Discount struct {
	// Currency code (e.g. BTC, ETH, USDT, etc.)
	Currency string `json:"currency"`
	// Blockchain network (e.g. ARBITRUM, AVALANCHE, BCH, etc.)
	Network string `json:"network"`
	// Discount percent (e.g. 5, 10, -15, etc.)
	Discount int `json:"discount"`
}

type ListOfDiscount []Discount

type ListOfDiscountResponse struct {
	*HTTPResponse
	Result ListOfDiscount `json:"result"`
}

// ListOfDiscount returns a list of discounts.
//
// Details: https://doc.cryptomus.com/business/discount/list
//
// Example:
//
//	result, err := sdk.ListOfDiscount()
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) ListOfDiscount() (*ListOfDiscountResponse, error) {
	return sdk.ListOfDiscountWithContext(context.Background())
}

// ListOfDiscountWithContext returns a list of discounts.
//
// Details: https://doc.cryptomus.com/business/discount/list
//
// Example:
//
//	result, err := sdk.ListOfDiscountWithContext(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) ListOfDiscountWithContext(ctx context.Context) (*ListOfDiscountResponse, error) {

	var result ListOfDiscountResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, nil)).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Post(ListOfDiscountsEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}

type SetDiscountToPaymentMethodRequest struct {
	// Currency code (e.g. BTC, ETH, USDT, etc.)
	Currency string `json:"currency"`
	// Blockchain network (e.g. ARBITRUM, AVALANCHE, BCH, etc.)
	Network string `json:"network"`
	// Discount percent (e.g. 5, 10, -15, etc.)
	Discount int `json:"discount_percent"`
}

type SetDiscountToPaymentMethodResponse struct {
	*HTTPResponse
	Result Discount `json:"result,omitempty"`
}

// SetDiscountToPaymentMethod sets a discount to a payment method.
//
// Details: https://doc.cryptomus.com/business/discount/set
//
// Discount Percent:
//   - Positive Numbers (>0): Provides buyers a discount for paying with a coin.
//     Useful for promotions or encouraging payments with a particular cryptocurrency.
//   - Negative Numbers (<0): Adds a percentage (padding) for paying with a coin.
//     This can help cover crypto/fiat conversion costs or adjust for local exchange differences.
//
// Example:
//
//	err := sdk.SetDiscountToPaymentMethod(payload)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) SetDiscountToPaymentMethod(payload SetDiscountToPaymentMethodRequest) (*SetDiscountToPaymentMethodResponse, error) {
	return sdk.SetDiscountToPaymentMethodWithContext(context.Background(), payload)
}

// SetDiscountToPaymentMethodWithContext sets a discount to a payment method.
//
// Details: https://doc.cryptomus.com/business/discount/set
//
// Discount Percent:
//   - Positive Numbers (>0): Provides buyers a discount for paying with a coin.
//     Useful for promotions or encouraging payments with a particular cryptocurrency.
//   - Negative Numbers (<0): Adds a percentage (padding) for paying with a coin.
//     This can help cover crypto/fiat conversion costs or adjust for local exchange differences.
//
// Example:
//
//	err := sdk.SetDiscountToPaymentMethodWithContext(ctx, payload)
//	if err != nil {
//	    log.Fatal(err)
//	}
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
