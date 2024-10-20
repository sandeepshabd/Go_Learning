/* 

The shared resource here is the balance field of the BankAccount struct. 
Two goroutines (go account.withdraw(700) and go account.withdraw(500)) are executed concurrently, both trying to withdraw funds from the same account. 
The race condition occurs in the withdraw method when both goroutines check the current balance (in the if statement) and proceed to withdraw money without proper synchronization. 
How the race unfolds: 

Goroutine 1 checks the balance (e.g., $1000) and sees that there is enough to withdraw $700. 
Goroutine 2 also checks the balance (still $1000, because Goroutine 1 hasn’t finished), and sees that there is enough to withdraw $500. 
Both goroutines then proceed to withdraw their amounts, unaware that the other is also withdrawing. 
The final balance could be incorrect, leading to both withdrawals being processed even though there wasn’t enough money in the account. 
Example of how the race occurs: 

Goroutine 1 sees balance = $1000, plans to withdraw $700. 
Goroutine 2 sees balance = $1000, plans to withdraw $500. 
Goroutine 1 withdraws $700, new balance = $300. 
Goroutine 2 also withdraws $500 based on the outdated balance, which should not have been allowed, leaving the account at an incorrect final balance of $-200. 
Without synchronization, the result could be unpredictable or lead to an invalid account balance, such as negative funds. 
 */



package main

import (
 "fmt"
 "sync"
 "time"
)

type BankAccount struct {
 balance int
}

func (a *BankAccount) withdraw(amount int) {
 if a.balance >= amount {
  // Simulate some processing delay
  time.Sleep(time.Millisecond * 100)
  a.balance -= amount
  fmt.Printf("Successfully withdrew $%d, remaining balance: $%d\n", amount, a.balance)
 } else {
  fmt.Printf("Failed to withdraw $%d, insufficient balance. Current balance: $%d\n", amount, a.balance)
 }
}

func main() {
 account := BankAccount{balance: 1000} // Initial balance of $1000
 var wg sync.WaitGroup

 wg.Add(2)

 // First withdrawal attempt
 go func() {
  defer wg.Done()
  account.withdraw(700) // Withdraw $700
 }()

 // Second withdrawal attempt
 go func() {
  defer wg.Done()
  account.withdraw(500) // Withdraw $500
 }()

 wg.Wait()
 fmt.Printf("Final balance: $%d\n", account.balance)
}