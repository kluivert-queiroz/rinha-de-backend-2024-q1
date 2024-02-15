package entities

import (
	"math"
	"time"
)

type Customer struct {
	Id           string        `bson:"_id" cql:"id"`
	Balance      int           `cql:"balance"`
	Limit        int           `cql:"limit"`
	Transactions []Transaction `cql:"transactions"`
	Version      int           `cql:"version"`
}

func NewCustomer(id string, balance int, limit int, transactions []Transaction) *Customer {
	c := Customer{
		Id:           id,
		Balance:      balance,
		Limit:        limit,
		Transactions: transactions,
	}
	return &c
}
func (c *Customer) Deposit(amount int, description string) {
	c.Balance += amount
	t := Transaction{
		Amount:      amount,
		Description: description,
		Type:        C,
		Date:        time.Now(),
	}
	c.Transactions = append([]Transaction{t}, c.Transactions...)
}

func (c *Customer) Withdraw(amount int, description string) {
	c.Balance -= amount
	t := Transaction{
		Amount:      amount,
		Description: description,
		Type:        D,
		Date:        time.Now(),
	}
	c.Transactions = append([]Transaction{t}, c.Transactions...)
	if len(c.Transactions) > 10 {
		c.Transactions = c.Transactions[0:10]
	}
}
func (c *Customer) Validate() error {
	if c.Balance > 0 {
		return nil
	}

	if int(math.Abs(float64(c.Balance))) > c.Limit {
		return ErrNotEnoughLimit
	}
	return nil
}
func (c *Customer) ID() string                        { return c.Id }
func (c *Customer) LatestTransactions() []Transaction { return c.Transactions }
