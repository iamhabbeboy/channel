package main

type IBankAccount interface {
	GetBalance() int
	Deposit(amount int)
	Withdraw(amount int) error
}
