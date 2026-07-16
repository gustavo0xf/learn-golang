package main

import (
	"testing"
)

func TestLoop(t *testing.T) {
	t.Run("testing counter", func(t *testing.T) {
		numbers := [10]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
		want := 4
		got := counter(numbers, 10)
		errHandling(t, got, want)
	})
	t.Run("testing sum", func(t *testing.T) {
		numbers := [10]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
		want := 30
		got := sum(numbers)
		errHandling(t, got, want)
	})
}

func BenchmarkLoop(b *testing.B) {
	b.Run("benchmarking counter", func(b *testing.B) {
		numbers := [10]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
		for b.Loop() {
			counter(numbers, 10)
		}
	})
	b.Run("benchmarking sum", func(b *testing.B) {
		numbers := [10]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
		for b.Loop() {
			sum(numbers)
		}
	})
}

func errHandling(t testing.TB, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
