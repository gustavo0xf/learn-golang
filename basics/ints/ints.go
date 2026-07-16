package main

import "fmt"

var (
	a int = 10
	b int = 20
)

func sum(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func mod(x int, y int) int {
	return x % y
}

func mul(x int, y int) int {
	return x * y
}

func main() {
	// short declaration (implicit) -> leva prioridade
	a := 30
	b := 40
	fmt.Println(sum(a, b))
	fmt.Println(sub(a, b))
	fmt.Println(mod(a, b))
	fmt.Println(mul(b, b))
}
