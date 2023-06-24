package utils_test

import (
	"learngowithtests/utils"
	"testing"
)

func TestEncodeJSON(t *testing.T) {
	data := map[string]string{
		"code":    "CAN_NOT_FIND_PRODUCT",
		"message": "cannot find product with the id",
	}

	want := `{"code":"CAN_NOT_FIND_PRODUCT","message":"cannot find product with the id"}`

	var got []byte
	got, err := utils.EncodeJSON(data)

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if string(got) != want {
		t.Errorf("want %s, got %s", want, got)
	}
}
