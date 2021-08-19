package account

import (
	"sync"
)

type Account struct {
	open    bool
	balance int64
	sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount, open: true}
}

func (ac *Account) Deposit(dep int64) (int64, bool) {
	ac.Mutex.Lock()
	defer ac.Mutex.Unlock()

	if !ac.open {
		return 0, false
	}
	if ac.balance+dep >= 0 {
		ac.balance += dep
		return ac.balance, true
	}
	return 0, false
}

func (ac *Account) Close() (int64, bool) {
	ac.Mutex.Lock()
	defer ac.Mutex.Unlock()

	var payout int64
	if ac.open {
		ac.open = false
		payout = ac.balance
		ac.balance = 0
		return payout, true
	}
	return payout, false
}

func (ac *Account) Balance() (int64, bool) {
	if ac.open {
		return ac.balance, true
	}
	return 0, false
}
