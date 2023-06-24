package interfacecasting_test

import (
	interfacecasting "learngowithtests/interface_casting"
	"learngowithtests/utils"
	"reflect"
	"testing"
)

func TestCastInterfaceAsError(t *testing.T) {

	apiErrorData := []*interfacecasting.Error{
		{
			Code:    "PRODUCT_OUT_OF_STOCK",
			Message: "product out of stock",
		},
		{
			Code:    "PRODUCT_OUT_OF_STOCK",
			Message: "product out of stock",
		},
	}

	apiErrorDataJSON, err := utils.EncodeJSON(apiErrorData)
	if err != nil {
		t.Fatalf("Unable to encode json - %v", err)
	}

	var apiErrorsInterface interface{}
	utils.DecodeJSON(apiErrorDataJSON, &apiErrorsInterface)

	got, err := interfacecasting.CastInterfaceAsError(apiErrorsInterface)

	if err != nil {
		t.Fatalf("Unexpected error - %v", err)
	}

	for i := 0; i < len(got); i++ {
		if !reflect.DeepEqual(got[i], apiErrorData[i]) {
			t.Errorf("want %#v, got %#v", *apiErrorData[i], *got[i])
		}
	}

}
