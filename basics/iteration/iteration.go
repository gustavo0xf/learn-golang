package main

import "fmt"

func counter(numbers [10]int, search int) int {
	counter := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] != search {
			counter++
		}
	}

	return counter
}

func sum(numbers [10]int) int {
	sum := 0
	// https://go.dev/doc/effective_go#blank
	// range -> retorna o indice e o valor do array. usamos o blank pra ignorar o indice
	for _, num := range numbers {
		sum += num
	}

	return sum
}

func main() {
	numbers := [10]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	fmt.Printf("[*] O(n) algorithm. take %d times to find element %d\n", counter(numbers, 10), 10)
	fmt.Printf("[*] array's elements sum = %d\n", sum(numbers))
}
