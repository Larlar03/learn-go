package main

import (
	"errors"
	"fmt"
)

type Simoleons int

type Wallet struct {
	balance Simoleons
}

type Stringer interface {
	String() string
}

func (s Simoleons) String() string {
	return fmt.Sprintf("§%d", s)
}

func (w *Wallet) Deposit(amount Simoleons) {
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("insufficient funds")

func (w *Wallet) Withdraw(amount Simoleons) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil

}

func (w *Wallet) Balance() Simoleons {
	return (*w).balance
}
