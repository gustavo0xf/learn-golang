package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}  // tamanho encodado no tipo, qlqr coisa ele nem compila
	slice := []int{2, 4, 6, 8, 10}  // flexivel quanto ao tamanho
	sliceTest := make([]int, 2, 10) // slice com capacidade e tamanho predefinidos
	fmt.Println(len(array))
	fmt.Println(len(slice))
	fmt.Println(cap(sliceTest))
}
