package main

import (
	"testing"
)

func TestMap(t *testing.T) {
	myMap := Map{"uuid": "89869fd7-7eaa-463f-987a-f224a7c85d36"}
	t.Run("test search", func(t *testing.T) {
		got, _ := myMap.searchMap("uuid")
		want := "89869fd7-7eaa-463f-987a-f224a7c85d36"
		errHandling(t, got, want)
	})
	t.Run("inexistent key", func(t *testing.T) {
		_, err := myMap.searchMap("test")
		want := "[!] unable to find the key on the map"
		if err == nil {
			t.Fatal("[!] expected to fail")
		}
		errHandling(t, err.Error(), want)
	})
	t.Run("test update", func(t *testing.T) {
		got, _ := myMap.updateMap("username", "ghu")
		want := "[*] map updated!"
		errHandling(t, got, want)
	})
	t.Run("key already exists", func(t *testing.T) {
		_, err := myMap.updateMap("uuid", "ghu")
		want := "[!] key already exists"
		if err == nil {
			t.Fatal("[!] expected to fail")
		}
		errHandling(t, err.Error(), want)
	})
}

func errHandling(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
