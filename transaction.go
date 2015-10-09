package main

import (
	"fmt"
	"sync"
)

type Fund struct {
	// balance is unexported (private), because it's lowercase
	balance int
}

// A regular function returning a pointer to a fund
func NewFund(initialBalance int) *Fund {
	// We can return a pointer to a new struct without worrying about
	// whether it's on the stack or heap: Go figures that out for us.
	return &Fund{
		balance: initialBalance,
	}
}

// Methods start with a *receiver*, in this case a Fund pointer
func (f *Fund) Balance() int {
	return f.balance
}

func (f *Fund) Withdraw(amount int) {
	f.balance -= amount
}

func main() {
	const WORKERS = 10
	val := 2000
	fund := NewFund(val)
	dollarsPerFounder := val / WORKERS
	fmt.Println("initial balance is", fund.Balance())

	var wg sync.WaitGroup

	for i := 0; i < WORKERS; i++ {
		// Let the waitgroup know we're adding a goroutine
		wg.Add(1)
		fmt.Println("worker", i+1, "initialized")

		// Spawn off a founder worker, as a closure
		go func() {
			// Mark this worker done when the function finishes
			defer wg.Done()

			for i := 0; i < dollarsPerFounder; i++ {
				fund.Withdraw(1)
			}

		}() // Remember to call the closure!
	}

	// Wait for all the workers to finish
	wg.Wait()
	fmt.Println("final balance is", fund.Balance())

}
