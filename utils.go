package cryptomus_sdk_go

import "encoding/json"

func ToJSON[T any](v T) ([]byte, error) {
	resByte, err := json.Marshal(v)
	if err != nil {
		return []byte{}, err
	}

	return resByte, nil
}

func FromJSONString[T any](v string) (T, error) {
	var res T
	err := json.Unmarshal([]byte(v), &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
