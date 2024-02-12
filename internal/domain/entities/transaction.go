package entities

import "time"

type Transaction struct {
	Amount      int       `json:"valor"`
	Description string    `json:"descricao"`
	Type        string    `json:"tipo"`
	Date        time.Time `json:"realizada_em"`
}
