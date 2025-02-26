package cryptomus

import "context"

type BalanceResponse struct {
	*HTTPResponse
	Result []BalanceResult `json:"result,omitempty"`
}

type Merchant struct {
	UUID         string `json:"uuid"`
	Balance      string `json:"balance"`
	CurrencyCode string `json:"currency_code"`
}

type User struct {
	UUID         string `json:"uuid"`
	Balance      string `json:"balance"`
	CurrencyCode string `json:"currency_code"`
}

type Balance struct {
	Merchant []Merchant `json:"merchant"`
	User     []User     `json:"user"`
}

type BalanceResult struct {
	Balance Balance `json:"balance"`
}

func (sdk *Cryptomus) Balance(ctx context.Context) (*BalanceResponse, error) {
	var result BalanceResponse

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetSuccessResult(&result).
		SetErrorResult(&result)

	if _, err := req.Get(BalanceEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}
