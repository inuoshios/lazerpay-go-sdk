package tests

import (
	"testing"

	"github.com/inuoshios/lazerpay-go-sdk"
)

func TestGetAcceptedCoins(t *testing.T) {
	coins, err := lazerpay.GetAcceptedCoins(clientHeader)

	if err != nil {
		// terrible error messages
		t.Errorf("could not get all coins -> [ERROR] %v -> [INFO] %s", err, coins)
	}
}
