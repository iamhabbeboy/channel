package main

import (
	"fmt"
	"math/rand"
	"time"
)

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func main() {
	// dataChan := make(chan int)
	// go func() {
	// 	wg := sync.WaitGroup{}
	// 	for i := 0; i < 1000; i++ {
	// 		wg.Add(1)
	// 		go func() {
	// 			defer wg.Done()
	// 			dataChan <- DoWork()
	// 		}()
	// 	}
	// 	wg.Wait()
	// 	close(dataChan)
	// }()

	// for n := range dataChan {
	// 	fmt.Printf("n = %d\n", n)
	// }

	// Context
	// DoContext()

	// Implementing interface
	// var bank BankAccount = Gtb{NewGtb().Deposit()}
	// g := NewGtb()
	// g.Deposit(200)
	// g.Deposit(100)
	// g.Withdraw(50)
	// if err := g.Withdraw(250); err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("The balance is %d", g.GetBalance())
	myAccounts := []IBankAccount{
		NewGtb(),
		NewBitcoinAccount(),
	}
	for _, account := range myAccounts {
		account.Deposit(200)
		account.Withdraw(50)
		balance := account.GetBalance()
		fmt.Printf("balance= %d\n", balance)
	}
}
