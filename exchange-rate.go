package cryptomus_sdk_go

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

func (sdk *Cryptomus) ExchangeRateList(ctx context.Context, currency string) (*ExchangeRateResponse, error) {
	url := fmt.Sprintf(ExchangeRateListEndpoint.URL(), currency)

	var result ExchangeRateResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Get(url); err != nil {
		return nil, err
	}

	return &result, nil
}
