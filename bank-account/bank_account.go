package account

import "sync"

// Account represents a bank account.
type Account struct {
	balance int64
	closed  bool
	mux     sync.Mutex
}

// Open opens an account.
func Open(deposit int64) *Account {
	if deposit < 0 {
		return nil
	}
	return &Account{deposit, false, sync.Mutex{}}
}

// Close closes an account.
func (a *Account) Close() (payout int64, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if a.closed {
		return 0, false
	}

	a.closed = true
	return a.balance, true
}

// Balance returns the current balance.
func (a *Account) Balance() (balance int64, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if a.closed {
		return 0, false
	}
	return a.balance, true
}

// Deposit handles a deposit to (or withdrawal from) the account.
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if a.closed {
		return 0, false
	}
	if a.balance+amount < 0 {
		return a.balance, false
	}
	a.balance = a.balance + amount
	return a.balance, true
}
