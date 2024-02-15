package entities

import (
	"fmt"
	"time"
)

const (
	D TransactionType = "d"
	C TransactionType = "c"
)

type (
	TransactionType string
	Transaction     struct {
		Amount      int             `json:"valor" cql:"amount"`
		Description string          `json:"descricao" cql:"description"`
		Type        TransactionType `json:"tipo" cql:"type"`
		Date        time.Time       `json:"realizada_em" cql:"date"`
	}
)

func (t TransactionType) Validate() error {
	switch t {
	case D:
	case C:
	default:
		return fmt.Errorf("invalid transaction type")
	}
	return nil
}
