package main

import "testing"

func TestOperations(t *testing.T) {
	t.Run("testing sum", func(t *testing.T) {
		want := 30
		got := sum(20, 10)
		errHandling(t, got, want)
	})
	t.Run("testing sub", func(t *testing.T) {
		want := 10
		got := sub(20, 10)
		errHandling(t, got, want)
	})
	t.Run("testing mod", func(t *testing.T) {
		want := 0
		got := mod(20, 10)
		errHandling(t, got, want)
	})
	t.Run("testing mul", func(t *testing.T) {
		want := 200
		got := mul(20, 10)
		errHandling(t, got, want)
	})
}

func errHandling(t testing.TB, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
