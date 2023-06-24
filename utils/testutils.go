package utils

import "encoding/json"

func EncodeJSON(data interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func DecodeJSON(jsonData []byte, v interface{}) error {
	err := json.Unmarshal(jsonData, v)
	if err != nil {
		return err
	}
	return nil
}
