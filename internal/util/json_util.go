package util

import "encoding/json"

func ParseJSONResponse[T any](jsonResponse []byte) (*T, error) {
	var data T
	err := json.Unmarshal(jsonResponse, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
