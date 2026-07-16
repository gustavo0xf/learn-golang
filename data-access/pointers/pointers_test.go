package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.depositBTC(10)
	got := wallet.getBalance()
	fmt.Printf("[!] btcAmount addr in test: %p\n", &wallet.btcAmount)
	want := 10
	errHandling(t, got, want)
}

func errHandling(t testing.TB, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
