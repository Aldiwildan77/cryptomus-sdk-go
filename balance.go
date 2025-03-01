package cryptomus

import "context"

type BalanceResponse struct {
	*HTTPResponse
	Result []BalanceResult `json:"result,omitempty"`
}

// Merchant represents a business account.
type Merchant struct {
	// Wallet UUID
	UUID string `json:"uuid"`
	// Business wallet balance
	Balance string `json:"balance"`
	// Wallet Currency code (e.g. BTC, ETH, USDT, etc.)
	CurrencyCode string `json:"currency_code"`
}

// User represents a personal account.
type User struct {
	// Wallet UUID
	UUID string `json:"uuid"`
	// Personal wallet balance
	Balance string `json:"balance"`
	// Wallet Currency code (e.g. BTC, ETH, USDT, etc.)
	CurrencyCode string `json:"currency_code"`
}

// Balance represents a list of balances.
//
// The list of balances includes Business (merchant) and Personal (user) wallets.
type Balance struct {
	Merchant []Merchant `json:"merchant"`
	User     []User     `json:"user"`
}

type BalanceResult struct {
	Balance Balance `json:"balance"`
}

// Balance returns a list of balances.
//
// Details: https://doc.cryptomus.com/business/balance
//
// Example:
//
//	result, err := sdk.Balance()
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) Balance() (*BalanceResponse, error) {
	return sdk.BalanceWithContext(context.Background())
}

// BalanceWithContext returns a list of balances.
//
// Details: https://doc.cryptomus.com/business/balance
//
// Example:
//
//	result, err := sdk.BalanceWithContext(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (sdk *Cryptomus) BalanceWithContext(ctx context.Context) (*BalanceResponse, error) {
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
