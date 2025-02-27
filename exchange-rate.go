package cryptomus

import (
	"context"
	"fmt"
)

type ExchangeRate struct {
	From   string `json:"from,omitempty"`
	To     string `json:"to,omitempty"`
	Course string `json:"course,omitempty"`
}

type ExchangeRateList []ExchangeRate

type ExchangeRateResponse struct {
	*HTTPResponse
	Result ExchangeRateList `json:"result,omitempty"`
}

func (sdk *Cryptomus) ExchangeRateList(currency string) (*ExchangeRateResponse, error) {
	return sdk.ExchangeRateListWithContext(context.Background(), currency)
}

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
