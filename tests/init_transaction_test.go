package tests

import (
	"testing"

	"github.com/inuoshios/lazerpay-go-sdk"
)

func TestInitTransaction(t *testing.T) {
	initialize := lazerpay.InitializeTransaction{
		Reference:            "4tytytreytrey65756u5u66",
		Amount:               "1000",
		CustomerName:         "Abdulfatai Suleiman",
		CustomerEmail:        "staticdev20046@gmail.com",
		Currency:             "USD",
		Coin:                 "DAI",
		AcceptPartialPayment: true,
	}

	res, err := lazerpay.InitTransaction(clientHeader, initialize)
	if err != nil {
		t.Errorf("[ERROR] -> Failed to initialize transaction: %v -> [INFO] -> %v", err, res)
	}
}
