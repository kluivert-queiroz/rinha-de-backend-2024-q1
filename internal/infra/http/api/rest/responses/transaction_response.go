package responses

type TransactionResponse struct {
	Balance int `json:"saldo"`
	Limit   int `json:"limite"`
}
