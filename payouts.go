package cryptomus_sdk_go

import "context"

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

type CreatePayoutData struct {
	UUID          string        `json:"uuid"`
	Amount        string        `json:"amount"`
	Currency      string        `json:"currency"`
	Network       string        `json:"network"`
	Address       string        `json:"address"`
	TxID          string        `json:"txid"`
	Status        PaymentStatus `json:"status"`
	IsFinal       bool          `json:"is_final"`
	Balance       int64         `json:"balance"`
	PayerCurrency string        `json:"payer_currency"`
	PayerAmount   int64         `json:"payer_amount"`
}

type CreatePayoutResponse struct {
	*HTTPResponse
	Result *CreatePayoutData `json:"result"`
}

func (sdk *Cryptomus) CreatePayout(ctx context.Context, payload *CreatePayoutRequest) (*CreatePayoutResponse, error) {
	var result CreatePayoutResponse

	payloadByte, err := ToJSON(payload)
	if err != nil {
		return nil, err
	}

	req := sdk.HttpClient.NewRequest().
		SetContext(ctx).
		SetHeader("merchant", sdk.Merchant).
		SetHeader("sign", Sign(sdk.PaymentToken, string(payloadByte))).
		SetHeader("Content-Type", "application/json").
		SetSuccessResult(&result).
		SetErrorResult(&result).
		SetBody(payloadByte)

	if _, err := req.Post(CreatePayoutEndpoint.URL()); err != nil {
		return nil, err
	}

	return &result, nil
}
