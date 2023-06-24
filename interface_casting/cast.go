package interfacecasting

import (
	"errors"
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// json -> decode by {}interface
// 1. {}interface -> []Error: cannot do this way, ok return false
// 2. {}interface -> []interface{} -> map[string]interface{} -> []Error: can do

func CastInterfaceAsError(err interface{}) ([]*Error, error) {
	interfaceSlice, ok := err.([]interface{})
	if !ok {
		return nil, errors.New("unable to cast interface{} as []interface{}")
	}

	var outErrors []*Error
	for _, item := range interfaceSlice {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, errors.New("unable to cast interface{} as Error")
		}

		code, _ := itemMap["code"].(string)
		message, _ := itemMap["message"].(string)

		outError := Error{
			Code:    code,
			Message: message,
		}

		outErrors = append(outErrors, &outError)
	}
	return outErrors, nil
}
