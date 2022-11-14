package tests

import (
	"testing"

	"github.com/inuoshios/lazerpay-go-sdk"
)

func TestVerifyTransaction(t *testing.T) {
	verify, err := lazerpay.VerifyTransaction(clientHeader, "0xf2345527195C3bdc6C5f07576a3C860281926841")

	if err != nil {
		t.Errorf("[INFO] -> unable to verify transaction with ref of %s [ERROR] %v", verify, err)
	}
}
