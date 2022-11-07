package tests

import (
	"github.com/milinches/lazerpay-go-sdk"
	"testing"
)

func TestGetAcceptedCoins(t *testing.T) {
	coins, err := lazerpay.GetAcceptedCoins(clientHeader)

	if err != nil {
		// terrible error messages
		t.Errorf("could not get all coins -> [ERROR] %v -> [INFO] %s", err, coins)
	}
}
