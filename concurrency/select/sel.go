package main

import (
	"fmt"
	"net/http"
)

func AppRacer(app1, app2 string) string {
	select {
	// blocking calls, i want the app that arrives first.
	case <-req(app1):
		return app1
	case <-req(app2):
		return app2
	}
}

func req(url string) chan struct{} {
	ch := make(chan struct{}) // use channel as a sruct to not allocate nothing!
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}

func main() {
	app1 := "https://google.com"
	app2 := "https://b4sh0xf.github.io"
	fmt.Println(AppRacer(app1, app2))
}
