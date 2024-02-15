package rest

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/avast/retry-go"
	goVal "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/application/services"
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/domain/entities"
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/infra/db/repositories"
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/infra/http/api/rest/requests"
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/infra/http/api/rest/responses"
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/pkg/validator"
)

type CustomerController struct {
	transactionService services.TransactionService
	validator          validator.XValidator
}

func NewCustomerController(t *services.TransactionService) *CustomerController {
	v := validator.NewValidator()
	v.Validator().RegisterValidation("transactionType", func(fl goVal.FieldLevel) bool {
		t := fl.Field().String()
		if t == "c" || t == "d" {
			return true
		}
		return false
	})
	return &CustomerController{transactionService: *t, validator: *v}
}
func (c *CustomerController) Transaction(ctx *fiber.Ctx) error {
	request := new(requests.TransactionRequest)
	ctx.BodyParser(request)

	request.Type = strings.ToLower(request.Type)

	if errs := c.validator.Validate(request); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"Field [%s] with value '%v' failed. Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		return &fiber.Error{
			Code:    fiber.ErrUnprocessableEntity.Code,
			Message: strings.Join(errMsgs, " and "),
		}
	}
	userId, paramErr := ctx.ParamsInt("id")
	if paramErr != nil {
		return &fiber.Error{
			Code:    fiber.ErrUnprocessableEntity.Code,
			Message: "Invalid customer id",
		}
	}
	var customer entities.Customer
	var err error
	retryableDepositOrWithdraw(func() error {
		if request.Type == string(entities.C) {
			customer, err = c.transactionService.Deposit(ctx.Context(), *request.ToDepositCommand(strconv.Itoa(userId)))
		}
		if request.Type == string(entities.D) {
			customer, err = c.transactionService.Withdraw(ctx.Context(), *request.ToWithdrawCommand(strconv.Itoa(userId)))
		}
		return err
	})
	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrUnprocessableEntity.Code,
			Message: err.Error(),
		}
	}
	return ctx.JSON(responses.TransactionResponse{Balance: customer.Balance, Limit: customer.Limit})
}

func (c *CustomerController) Statements(ctx *fiber.Ctx) error {
	userId, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(400).JSON(map[string]string{"reason": "Invalid customer id"})
	}
	customer, err := c.transactionService.RetrieveCustomer(ctx.Context(), strconv.Itoa(userId))
	if err != nil {
		if errors.Is(err, repositories.ErrCustomerNotFound) {
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

func retryableDepositOrWithdraw(fn func() error) error {
	configs := []retry.Option{
		retry.Attempts(uint(10)),
		retry.OnRetry(func(n uint, err error) {
			log.Infof("Retry request %d to and get error: %v", n+1, err)
		}),
	}
	err := retry.Do(
		fn, configs...,
	)
	return err
}
