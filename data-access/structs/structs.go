package main

import (
	"fmt"
)

type Customer struct {
	id       int
	balance  float64
	username string
	password string
}

func getId(customer Customer) int {
	return customer.id
}

func getName(customer Customer) string {
	return customer.username
}

func main() {
	customer := Customer{1, 100.0, "ghu", "vasco"}
	fmt.Println(getId(customer))
	fmt.Println(getName(customer))
}
