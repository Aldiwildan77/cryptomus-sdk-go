package cryptomus

import "encoding/json"

// ToJSON converts a struct to a JSON byte array.
//
// Example:
//
//	type ExchangeRate struct {
//		From   string `json:"from,omitempty"`
//		To     string `json:"to,omitempty"`
//		Course string `json:"course,omitempty"`
//	}
//
//	exchangeRate := ExchangeRate{
//		From:   "ETH",
//		To:     "BTC",
//		Course: "0.0001",
//	}
//
//	resByte, err := ToJSON[ExchangeRate](exchangeRate)
//	if err != nil {
//		log.Fatal(err)
//	}
func ToJSON[T any](v T) ([]byte, error) {
	resByte, err := json.Marshal(v)
	if err != nil {
		return []byte{}, err
	}

	return resByte, nil
}

// FromJSONString converts a JSON string to a struct.
//
// Example:
//
//	type ExchangeRate struct {
//		From   string `json:"from,omitempty"`
//		To     string `json:"to,omitempty"`
//		Course string `json:"course,omitempty"`
//	}
//
//	exchangeRate, err := FromJSONString[ExchangeRate](`{"from":"ETH","to":"BTC","course":"0.0001"}`)
//	if err != nil {
//		log.Fatal(err)
//	}
func FromJSONString[T any](v string) (T, error) {
	var res T
	err := json.Unmarshal([]byte(v), &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
