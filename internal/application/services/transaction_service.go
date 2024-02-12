package services

import (
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/application/commands"
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/domain"
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/domain/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionService struct {
	customerRepository domain.CustomerRepository
}

func NewTransactionService(customerRepository domain.CustomerRepository) *TransactionService {
	return &TransactionService{
		customerRepository: customerRepository,
	}
}

func (t *TransactionService) Deposit(c commands.DepositCommand) (entities.Customer, error) {
	customer, err := t.customerRepository.FindById(c.CustomerId)
	if err != nil {
		return entities.Customer{}, err
	}
	customer.Deposit(c.Amount, c.Description)
	if err = t.customerRepository.Save(customer); err != nil {
		return entities.Customer{}, err
	}
	return customer, nil
}
func (t *TransactionService) Withdraw(c commands.WithdrawCommand) (entities.Customer, error) {
	customer, err := t.customerRepository.FindById(c.CustomerId)
	if err != nil {
		return entities.Customer{}, err
	}
	customer.Withdraw(c.Amount, c.Description)
	if err = customer.Validate(); err != nil {
		return entities.Customer{}, err
	}
	if err = t.customerRepository.Save(customer); err != nil {
		return entities.Customer{}, err
	}
	return customer, nil
}

func (t *TransactionService) RetrieveCustomer(customerId int) (entities.Customer, error) {
	customer, err := t.customerRepository.FindById(customerId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return entities.Customer{}, ErrCustomerNotFound
		}
		return entities.Customer{}, err
	}
	return customer, nil
}
