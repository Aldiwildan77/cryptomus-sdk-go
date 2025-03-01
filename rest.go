package cryptomus

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
)

type Errors map[string][]string

type HTTPResponse struct {
	State   int    `json:"state,omitempty"`
	Code    int    `json:"code,omitempty"`
	Errors  Errors `json:"errors,omitempty"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

type Pagination struct {
	Count          int    `json:"count"`
	HasPages       bool   `json:"hasPages"`
	NextCursor     string `json:"nextCursor,omitempty"`
	PreviousCursor string `json:"previousCursor,omitempty"`
	PerPage        int    `json:"perPage"`
}

func Sign(apiKey string, data interface{}) string {
	combined := apiKey

	if data != nil && data != "" {
		body, _ := json.Marshal(data)
		encodedData := base64.StdEncoding.EncodeToString(body)
		combined = encodedData + apiKey
	}

	hash := md5.Sum([]byte(combined))
	return hex.EncodeToString(hash[:])
}
