package cryptomus

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
	cryptomus := DefaultCryptomus()

	for _, option := range options {
		option(cryptomus)
	}

	return cryptomus
}

func DefaultCryptomus() *Cryptomus {
	return &Cryptomus{
		HttpClient: DefaultHTTPClient(),
	}
}

func WithHttpClient(client *req.Client) Option {
	return func(c *Cryptomus) {
		c.HttpClient = client
	}
}

func WithMerchant(merchant string) Option {
	return func(c *Cryptomus) {
		c.Merchant = merchant
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
