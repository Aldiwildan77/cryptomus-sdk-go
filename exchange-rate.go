package cryptomus

import (
	"context"
	"fmt"
)

type ExchangeRate struct {
	// From Currency code (e.g. BTC, ETH, USDT, etc.)
	From string `json:"from,omitempty"`
	// To Currency code (e.g. BTC, ETH, USDT, etc.)
	To string `json:"to,omitempty"`
	// Exchange rate (e.g. 0.0001, 0.0002, 0.0003, etc.)
	Course string `json:"course,omitempty"`
}

type ExchangeRateList []ExchangeRate

type ExchangeRateResponse struct {
	*HTTPResponse
	Result ExchangeRateList `json:"result,omitempty"`
}

// ExchangeRateList returns a list of exchange rates.
//
// Details: https://doc.cryptomus.com/business/exchange-rates/list
//
// Example:
//
//	result, err := sdk.ExchangeRateList("ETH")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) ExchangeRateList(currency string) (*ExchangeRateResponse, error) {
	return sdk.ExchangeRateListWithContext(context.Background(), currency)
}

// ExchangeRateListWithContext returns a list of exchange rates.
//
// Details: https://doc.cryptomus.com/business/exchange-rates/list
//
// Example:
//
//	result, err := sdk.ExchangeRateListWithContext(ctx, "ETH")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) ExchangeRateListWithContext(ctx context.Context, currency string) (*ExchangeRateResponse, error) {
	url := fmt.Sprintf(ExchangeRateListEndpoint.URL(), currency)

	var result ExchangeRateResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Get(url); err != nil {
		return nil, err
	}

	return &result, nil
}
