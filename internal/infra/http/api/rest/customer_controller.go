package rest

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/application/services"
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/domain/entities"
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/infra/http/api/rest/requests"
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/infra/http/api/rest/responses"
)

type CustomerController struct {
	transactionService services.TransactionService
}

func NewCustomerController(t *services.TransactionService) *CustomerController {
	return &CustomerController{transactionService: *t}
}
func (c *CustomerController) Transaction(ctx *fiber.Ctx) error {
	request := new(requests.TransactionRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}
	userId, paramErr := ctx.ParamsInt("id")
	if paramErr != nil {
		return ctx.Status(400).JSON(map[string]string{"reason": "Invalid customer id"})
	}
	var customer entities.Customer
	var err error
	if request.Type == "C" {
		customer, err = c.transactionService.Deposit(*request.ToDepositCommand(userId))
	}
	if request.Type == "D" {
		customer, err = c.transactionService.Withdraw(*request.ToWithdrawCommand(userId))
	}
	if err != nil {
		return ctx.Status(422).JSON(map[string]string{"reason": err.Error()})
	}
	return ctx.JSON(responses.TransactionResponse{Balance: customer.Balance, Limit: customer.Limit})
}

func (c *CustomerController) Statements(ctx *fiber.Ctx) error {
	userId, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(map[string]string{"reason": "Invalid customer id"})
	}
	customer, err := c.transactionService.RetrieveCustomer(userId)
	if err != nil {
		if errors.Is(err, services.ErrCustomerNotFound) {
			return ctx.Status(404).JSON(map[string]string{"reason": err.Error()})
		}
		log.Error(err)
		return ctx.Status(500).JSON(map[string]string{"reason": "Internal server error"})
	}
	response := responses.StatementsResponse{
		Balance:    responses.BalanceResponse{Total: customer.Balance, Limit: customer.Limit, Date: time.Now()},
		Statements: customer.LatestTransactions(),
	}
	return ctx.Status(200).JSON(response)
}
