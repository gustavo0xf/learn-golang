package main

import (
	// "bytes" -> bytes.Buffer cannot be passed as an io.Writer
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "%s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "hello bcraff")
}

func main() {
	log.Fatal(http.ListenAndServe(":8081", http.HandlerFunc(MyGreeterHandler)))
}

/*
returns the number of bytes written and any write error encountered.
func Printf(format string, a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintf(format, a)
	n, err = w.Write(p.buf)
	p.free()
	return
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
*/
