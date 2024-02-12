package responses

import (
	"time"

	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/domain/entities"
)

type StatementsResponse struct {
	Balance    BalanceResponse        `json:"saldo"`
	Statements []entities.Transaction `json:"ultimas_transacoes"`
}

type BalanceResponse struct {
	Total int       `json:"total"`
	Date  time.Time `json:"data_extrato"`
	Limit int       `json:"limite"`
}
