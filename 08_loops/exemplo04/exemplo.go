package main

import "fmt"

func main() {
	fmt.Println("Digite um número inteiro positivo: ")
	var n int
	fmt.Scanf("%d", &n)

	if n < 0 {
		n *= -1
	} else if n == 0 {
		n = 1
	}

	for contador := n; contador >= 1; contador-- {
		fmt.Printf("%d ", contador)
	}
}
