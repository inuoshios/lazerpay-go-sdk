package tests

import (
	"testing"

	"github.com/inuoshios/lazerpay-go-sdk"
)

func TestTransfer(t *testing.T) {
	initTransfer := lazerpay.TransferFunds{
		Amount:     1,
		Recipient:  "0xF378c952d5266eF8e1783521a1395Fe40cDCe55B",
		Coin:       "USDT",
		BlockChain: "Binance Smart Chain",
	}

	transfer, err := lazerpay.Transfer(clientHeader, initTransfer)
	if err != nil {
		t.Errorf("[ERROR] -> %v cannot transfer funds [INFO] -> %s ", err, transfer)
	}
}
