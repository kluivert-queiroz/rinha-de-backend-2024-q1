package requests

import "github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/application/commands"

type TransactionRequest struct {
	Amount      int    `json:"valor"`
	Description string `json:"descricao"`
	Type        string `json:"tipo"`
}

func (t *TransactionRequest) ToDepositCommand(customerId int) *commands.DepositCommand {
	return &commands.DepositCommand{
		Amount:      t.Amount,
		Description: t.Description,
		CustomerId:  customerId,
	}
}

func (t *TransactionRequest) ToWithdrawCommand(customerId int) *commands.WithdrawCommand {
	return &commands.WithdrawCommand{
		Amount:      t.Amount,
		Description: t.Description,
		CustomerId:  customerId,
	}
}
