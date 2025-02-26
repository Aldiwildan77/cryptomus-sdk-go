package cryptomus

import (
	"time"

	"github.com/imroc/req/v3"
)

func DefaultHTTPClient() *req.Client {
	maxTimeout := 5 * time.Second
	userAgent := "Cryptomus SDK Go"

	return req.
		NewClient().
		SetTimeout(maxTimeout).
		SetUserAgent(userAgent).
		SetCommonHeader("Content-Type", "application/json").
		SetCommonHeader("X-SDK-Language", "go").
		EnableDumpAllAsync()
}
