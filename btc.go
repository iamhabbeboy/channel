package main

import "errors"

type BitcoinAccount struct {
	balance int
	fee     int
}

func NewBitcoinAccount() *BitcoinAccount {
	return &BitcoinAccount{
		balance: 0,
		fee:     20,
	}
}

func (g *BitcoinAccount) GetBalance() int {
	return g.balance
}

func (g *BitcoinAccount) Deposit(amount int) {
	g.balance += amount
}

func (g *BitcoinAccount) Withdraw(amount int) error {
	newBalance := g.balance - amount - g.fee
	if newBalance < 0 {
		return errors.New("Insufficient funds")
	}
	g.balance = newBalance
	return nil
}
