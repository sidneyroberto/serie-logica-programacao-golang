package main

import "fmt"

func main() {
	fmt.Println("Digite uma palavra:")
	var palavra string
	fmt.Scanf("%s", &palavra)

	tamanho := len(palavra)
	metade := tamanho / 2

	fmt.Printf("Primeira metade: %s\n", palavra[0:metade])
	fmt.Printf("Segunda metade: %s\n", palavra[metade:])
}
