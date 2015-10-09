//author siddhant deshmukh
//UFID: 36568046
package main

import (
	"fmt"
	"sync"
)

type Account struct {
	balance int
}

func NewAccount(initBal int) *Account {

	return &Account{
		balance: initBal,
	}
}

func (a *Account) Balance() int {
	return a.balance
}

func (a *Account) Withdraw(amt int) {
	a.balance -= amt
}

func main() {
	var numOfWorkers int
	fmt.Println("enter the number of workers")
	fmt.Scan(&numOfWorkers)

	var val int
	fmt.Println("enter the amount in the bank account")
	fmt.Scan(&val)

	Acc := NewAccount(val)
	moneyPerWorker := val / numOfWorkers
	fmt.Println("initial balance is", Acc.Balance())

	var waitmaster sync.WaitGroup

	for i := 0; i < numOfWorkers; i++ {

		waitmaster.Add(1)
		fmt.Println("worker number ", i+1, "initialized")

		go func() {

			defer waitmaster.Done()

			for k := 0; k < moneyPerWorker; k++ {
				Acc.Withdraw(1)

			}

		}()
	}

	waitmaster.Wait()
	fmt.Println("final balance is", Acc.Balance())

}
