package pointersanderrors

import "fmt"

type Wallet struct {
	balance int
}

// In Go, when you call a function or a method the arguments are copied.

func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	// *Wallet is called `struct pointers` and they are automatically dereferenced
	// by convention you should keep your method receiver types the same for consistency.

	return w.balance
}
