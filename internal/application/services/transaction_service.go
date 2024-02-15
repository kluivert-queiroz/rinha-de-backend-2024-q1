package services

import (
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/application/commands"
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/domain"
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/domain/entities"
	"github.com/valyala/fasthttp"
)

type TransactionService struct {
	customerRepository domain.CustomerRepository
}

func NewTransactionService(customerRepository domain.CustomerRepository) *TransactionService {
	return &TransactionService{
		customerRepository: customerRepository,
	}
}

func (t *TransactionService) Deposit(ctx *fasthttp.RequestCtx, c commands.DepositCommand) (entities.Customer, error) {
	customer, err := t.customerRepository.FindById(ctx, c.CustomerId)
	if err != nil {
		return entities.Customer{}, err
	}
	customer.Deposit(c.Amount, c.Description)
	if err = t.customerRepository.Save(ctx, customer); err != nil {
		return entities.Customer{}, err
	}
	return customer, nil
}
func (t *TransactionService) Withdraw(ctx *fasthttp.RequestCtx, c commands.WithdrawCommand) (entities.Customer, error) {
	customer, err := t.customerRepository.FindById(ctx, c.CustomerId)
	if err != nil {
		return entities.Customer{}, err
	}
	customer.Withdraw(c.Amount, c.Description)
	if err = customer.Validate(); err != nil {
		return entities.Customer{}, err
	}
	if err = t.customerRepository.Save(ctx, customer); err != nil {
		return entities.Customer{}, err
	}
	return customer, nil
}

func (t *TransactionService) RetrieveCustomer(ctx *fasthttp.RequestCtx, customerId string) (entities.Customer, error) {
	customer, err := t.customerRepository.FindById(ctx, customerId)
	if err != nil {
		return entities.Customer{}, err
	}
	return customer, nil
}
