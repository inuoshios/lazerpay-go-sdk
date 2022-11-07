package lazerpay

type InitializeTransaction struct {
	Reference            string `json:"reference"`
	Amount               string `json:"amount"`
	CustomerName         string `json:"customer_name"`
	CustomerEmail        string `json:"customer_email"`
	Currency             string `json:"currency"`
	Coin                 string `json:"coin"`
	AcceptPartialPayment bool   `json:"accept_partial_payment"`
}

type TransferFunds struct {
	Amount     uint   `json:"amount"`
	Recipient  string `json:"recipient"`
	Coin       string `json:"coin"`
	BlockChain string `json:"blockchain"`
}
