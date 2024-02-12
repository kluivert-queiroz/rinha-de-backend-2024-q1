package entities

import (
	"math"
	"time"
)

type Customer struct {
	Id           int `bson:"_id"`
	Balance      int
	Limit        int
	Transactions []Transaction
}

func NewCustomer(id int, balance int, limit int, transactions []Transaction) *Customer {
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
	c.Transactions = append(c.Transactions, Transaction{
		Amount:      amount,
		Description: description,
		Type:        "C",
		Date:        time.Now(),
	})
}

func (c *Customer) Withdraw(amount int, description string) {
	c.Balance -= amount
	c.Transactions = append(c.Transactions, Transaction{
		Amount:      amount,
		Description: description,
		Type:        "D",
		Date:        time.Now(),
	})
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
func (c *Customer) ID() int                           { return c.Id }
func (c *Customer) LatestTransactions() []Transaction { return c.Transactions }
