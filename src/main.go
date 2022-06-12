package main

import (
	"fmt"
	"sync"
	"time"
)

// Concurrency control using mutual exclusion strategy
// Case: Two entities are expected to withdraw money concurrently, thus causing
// the second goroutine to overwrite (lost update) the result of the first goroutine;

func main() {
	balance := 100
	m := sync.Mutex{}
	wg := sync.WaitGroup{}

	fmt.Println("Account balance", balance)

	wg.Add(2)
	// Withdraw $20
	go func(balance *int, m *sync.Mutex, wg *sync.WaitGroup) {
		m.Lock()
		newBalance := *balance - 20
		time.Sleep(100 * time.Millisecond)
		*balance = newBalance
		m.Unlock()
		wg.Done()
	}(&balance, &m, &wg)

	// Withdraw $70
	go func(balance *int, m *sync.Mutex, wg *sync.WaitGroup) {
		m.Lock()
		newBalance := *balance - 70
		time.Sleep(10 * time.Millisecond)
		*balance = newBalance
		m.Unlock()
		wg.Done()
	}(&balance, &m, &wg)

	wg.Wait()

	time.Sleep(100 * time.Millisecond)
	// Expected balance is $10 when concurrency is well handled
	// otherwise is $80
	fmt.Println("Account balance", balance)
}
