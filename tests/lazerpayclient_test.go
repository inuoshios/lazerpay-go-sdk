package tests

import (
	"testing"

	"github.com/inuoshios/lazerpay-go-sdk"
)

func TestLazerpayClient(t *testing.T) {
	testCase, err := lazerpay.LazerpayClient(publicKey, secretKey)

	if err != nil {
		t.Errorf("[ERROR] could not add headers %v", testCase)
	}
}
