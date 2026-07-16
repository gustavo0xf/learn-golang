package main

type addrPing func(string) bool
type result struct {
	string
	bool
}

func CheckAddress(check addrPing, addrs []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, addr := range addrs {
		// every single call to check() starts on a new goroutine (concurrency!)
		go func() {
			resultChannel <- result{addr, check(addr)}
		}()
	}

	for i := 0; i < len(addrs); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

/*
-> takes too long:
for _, addr := range addrs {
	results[addr] = check(addr)
}

-> returns a empty map, because the CheckAddress returns fastly than the goroutines
for _, addr := range addrs {
	go func() {
		results[addr] = check(addr)
	}()
}

-> add a delay of 2 seconds: tests are inconclusive
	- test 1: ok      learn/concurrency       2.009s
	- test 2: fatal error: concurrent map writes
		- race condition, the map recieves two concomitant tries do write
		- two operations tried to write on the same memory address at the same time
		WARNING: DATA RACE
		Write at 0x00c0000a0288 by goroutine 92:
		learn/concurrency.CheckAddress.func1()
			/mnt/c/Users/TRT22/Downloads/learn-golang/concurrency/conc.go:13 +0x84

		Previous write at 0x00c0000a0288 by goroutine 29:
		learn/concurrency.CheckAddress.func1()
			/mnt/c/Users/TRT22/Downloads/learn-golang/concurrency/conc.go:13 +0x84

-> solution: use channels to coordinate the goroutines
	- comparsion:

*/
