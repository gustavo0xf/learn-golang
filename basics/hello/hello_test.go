package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello in specified language", func(t *testing.T) {
		got := hello("ghu", "portuguese")
		want := "ola, ghu"
		errHandling(t, got, want)
	})

	t.Run("all empty = troll", func(t *testing.T) {
		got := hello("", "")
		want := "troll"
		errHandling(t, got, want)
	})
}

func errHandling(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

/*

[*] testing rules

- It needs to be in a file with a name like xxx_test.go
- The test function must start with the word Test
- The test function takes one argument only t *testing.T
- To use the *testing.T type, you need to import "testing", like we did with fmt in the other file

*/
