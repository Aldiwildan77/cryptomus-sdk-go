package cryptomus

import (
	"time"

	"github.com/imroc/req/v3"
)

type Cryptomus struct {
	HttpClient   *req.Client
	Merchant     string
	PaymentToken string
	PayoutToken  string
}

type Option func(*Cryptomus)

func New(options ...Option) *Cryptomus {
	cryptomus := DefaultCryptomus()

	for _, option := range options {
		option(cryptomus)
	}

	return cryptomus
}

func DefaultCryptomus() *Cryptomus {
	maxTimeout := 10 * time.Second
	userAgent := "Cryptomus SDK Go"

	httpClient := req.
		NewClient().
		SetTimeout(maxTimeout).
		SetUserAgent(userAgent).
		SetCommonHeader("Content-Type", "application/json").
		SetCommonHeader("X-SDK-Language", "go").
		EnableDumpAllAsync()

	return New(
		WithHttpClient(httpClient),
	)
}

func WithHttpClient(client *req.Client) Option {
	return func(c *Cryptomus) {
		c.HttpClient = client
	}
}

func WithMerchantID(merchantID string) Option {
	return func(c *Cryptomus) {
		c.Merchant = merchantID
	}
}

func WithPaymentToken(token string) Option {
	return func(c *Cryptomus) {
		c.PaymentToken = token
	}
}

func WithPayoutToken(token string) Option {
	return func(c *Cryptomus) {
		c.PayoutToken = token
	}
}
