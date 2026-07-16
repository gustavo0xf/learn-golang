package main

import "testing"

func TestAppRacer(t *testing.T) {
	t.Run("racing", func(t *testing.T) {
		//app1 := "https://ghu.com"
		//app2 := "https://bcraff.com"
		//got, _ := AppRacer(app1, app2)

		//errHandling(t, got, app2)
	})
}

func errHandling(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
