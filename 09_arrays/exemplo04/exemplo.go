package main

import "fmt"

func main() {
	numeros := [5]int{1, 2, 3, 4, 5}

	soma := 0
	for i := 0; i < len(numeros); i++ {
		soma += numeros[i]
	}

	fmt.Printf("Somatório: %d\n", soma)
}
