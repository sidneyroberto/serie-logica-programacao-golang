package main

import "fmt"

func main() {
	var pessoas []string
	tamanho := len(pessoas)
	capacidade := cap(pessoas)
	fmt.Printf("Tamanho: %d Capacidade: %d Pessoas: %v\n", tamanho, capacidade, pessoas)

	pessoas = append(pessoas, "Sidney")
	tamanho = len(pessoas)
	capacidade = cap(pessoas)
	fmt.Printf("Tamanho: %d Capacidade: %d Pessoas: %v\n", tamanho, capacidade, pessoas)

	pessoas = append(pessoas, "Suelen")
	tamanho = len(pessoas)
	capacidade = cap(pessoas)
	fmt.Printf("Tamanho: %d Capacidade: %d Pessoas: %v\n", tamanho, capacidade, pessoas)

	pessoas = pessoas[:1]
	tamanho = len(pessoas)
	capacidade = cap(pessoas)
	fmt.Printf("Tamanho: %d Capacidade: %d Pessoas: %v\n", tamanho, capacidade, pessoas)

	pessoas = pessoas[:2]
	tamanho = len(pessoas)
	capacidade = cap(pessoas)
	fmt.Printf("Tamanho: %d Capacidade: %d Pessoas: %v\n", tamanho, capacidade, pessoas)
}
