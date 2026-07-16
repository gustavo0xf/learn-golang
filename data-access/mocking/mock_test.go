package main

import (
	"bytes"
	"testing"
)

type hookSleeper struct {
	calls int
}

func (h *hookSleeper) Sleep() {
	h.calls++
}

func TestCountdown(t *testing.T) {
	t.Run("count", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		h := &hookSleeper{}
		startCountdown(buffer, h)
		want := "3\n2\n1\ngo!"
		errHandling(t, buffer.String(), want)
		checkHook(t, h.calls)
	})
}

func errHandling(t testing.TB, got any, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func checkHook(t testing.TB, calls int) {
	t.Helper()
	if calls != 3 {
		t.Errorf("[!] not enough calls to sleep (%d)", calls)
	}
}
