package main

import "errors"

type Gtb struct {
	balance int
}

func NewGtb() *Gtb {
	return &Gtb{
		balance: 0,
	}
}

func (g *Gtb) GetBalance() int {
	return g.balance
}

func (g *Gtb) Deposit(amount int) {
	g.balance += amount
}

func (g *Gtb) Withdraw(amount int) error {
	newBalance := g.balance - amount
	if newBalance < 0 {
		return errors.New("Insufficient funds")
	}
	g.balance = newBalance
	return nil
}
