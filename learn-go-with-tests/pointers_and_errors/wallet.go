package wallet

import (
	"fmt"
	"errors"
)

type Bitcoin float64

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Wallet struct {
	balance Bitcoin // if a symbol (variables, types, functions et al) starts with a lowercase symbol then it is private outside the package it's defined in.
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
    return fmt.Sprintf("%g BTC", b)
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance // == (*w).total
}

func (w *Wallet) Deposit(value Bitcoin) {
	w.balance += value
}

func (w *Wallet) Withdraw(value Bitcoin) error {
    if value > w.balance {
        return ErrInsufficientFunds
    }
	w.balance -= value
	return nil
}