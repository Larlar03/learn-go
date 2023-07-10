package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Simoleons(10))

		want := Simoleons(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Simoleons(20)}

		wallet.Withdraw(Simoleons(10))

		want := Simoleons(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Simoleons(20)}

		err := wallet.Withdraw(Simoleons(100))

		assertError(t, err, ErrInsufficientFunds.Error())
		assertBalance(t, wallet, Simoleons(20))
	})

}

func assertBalance(t testing.TB, wallet Wallet, want Simoleons) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
