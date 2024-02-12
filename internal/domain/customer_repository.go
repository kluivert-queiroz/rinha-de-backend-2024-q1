package domain

import "github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/domain/entities"

type CustomerRepository interface {
	FindById(id int) (entities.Customer, error)
	Save(c entities.Customer) error
}
