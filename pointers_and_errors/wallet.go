package pointersanderrors

type Wallet struct {
	balance int
}

// In Go, when you call a function or a method the arguments are copied.

func (w Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}
