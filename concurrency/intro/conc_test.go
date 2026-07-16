package main

import (
	"testing"
	"time"
)

// simular o delay
func slowStubWebsiteChecker(string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckAddress(b *testing.B) {
	addrs := make([]string, 100)
	for i := 0; i < len(addrs); i++ {
		addrs[i] = "127.0.0.1"
	}
	for b.Loop() {
		CheckAddress(slowStubWebsiteChecker, addrs)
	}
}
