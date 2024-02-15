package requests

import "github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/application/commands"

type TransactionRequest struct {
	Amount      int    `json:"valor" validate:"required,numeric,min=1"`
	Description string `json:"descricao" validate:"required,min=1,max=10"`
	Type        string `json:"tipo" validate:"required,transactionType"`
}

func (t *TransactionRequest) ToDepositCommand(customerId string) *commands.DepositCommand {
	return &commands.DepositCommand{
		Amount:      t.Amount,
		Description: t.Description,
		CustomerId:  customerId,
	}
}

func (t *TransactionRequest) ToWithdrawCommand(customerId string) *commands.WithdrawCommand {
	return &commands.WithdrawCommand{
		Amount:      t.Amount,
		Description: t.Description,
		CustomerId:  customerId,
	}
}
