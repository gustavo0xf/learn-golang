package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("test msg", func(t *testing.T) {
		buffer := bytes.Buffer{}
		//Hello(&buffer, "hello bcraff")
		errHandling(t, buffer.String(), "hello bcraff")
	})
}

func errHandling(t testing.TB, got any, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
