package main

import (
	// "bytes"
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type defaultSleeper struct{}

func (d *defaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func startCountdown(w io.Writer, s Sleeper) {
	for i := 3; i > 0; i-- {
		s.Sleep()
	}
	for i := 3; i > 0; i-- {
		fmt.Fprintln(w, i)
	}
	fmt.Fprint(w, "go!")
}

func main() {
	sleeper := &defaultSleeper{}
	startCountdown(os.Stdout, sleeper)
}
