package main

import "fmt"

func main() {
	numeros := make([]int, 5)
	tamanho := len(numeros)
	capacidade := cap(numeros)
	fmt.Printf("Tamanho: %d Capacidade: %d Numeros: %v\n", tamanho, capacidade, numeros)

	numeros = make([]int, 5, 10)
	tamanho = len(numeros)
	capacidade = cap(numeros)
	fmt.Printf("Tamanho: %d Capacidade: %d Numeros: %v\n", tamanho, capacidade, numeros)

	numeros = numeros[:10]
	tamanho = len(numeros)
	capacidade = cap(numeros)
	fmt.Printf("Tamanho: %d Capacidade: %d Numeros: %v\n", tamanho, capacidade, numeros)
}
