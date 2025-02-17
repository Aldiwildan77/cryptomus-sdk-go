package cryptomus_sdk_go

import (
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
	cryptomus := &Cryptomus{
		HttpClient: req.NewClient(),
	}

	for _, option := range options {
		option(cryptomus)
	}

	return cryptomus
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
