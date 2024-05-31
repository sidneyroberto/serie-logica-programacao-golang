package main

import "fmt"

func main() {
	frutas := []string{"banana", "maçã", "uva", "goiaba"}
	tamanho := len(frutas)
	capacidade := cap(frutas)
	fmt.Printf("Tamanho: %d Capacidade: %d Frutas: %v\n", tamanho, capacidade, frutas)

	fatia := frutas[:2]
	tamanho = len(fatia)
	capacidade = cap(fatia)
	fmt.Printf("Tamanho: %d Capacidade: %d Fatia: %v\n", tamanho, capacidade, fatia)

	fatia = frutas[:3]
	tamanho = len(fatia)
	capacidade = cap(fatia)
	fmt.Printf("Tamanho: %d Capacidade: %d Fatia: %v\n", tamanho, capacidade, fatia)

	fatia = frutas[2:]
	tamanho = len(fatia)
	capacidade = cap(fatia)
	fmt.Printf("Tamanho: %d Capacidade: %d Fatia: %v\n", tamanho, capacidade, fatia)
}
