package main

import "fmt"

func main() {
	fmt.Println("Digite o seu nome:")
	var nome string        // Declara a variável nome
	fmt.Scanf("%s", &nome) // Armazena o nome do usuário na variável nome

	serie := "Lógica de Programação com Go"
	fmt.Printf("Seja bem-vindo(a) à série %s, %s!", serie, nome)
}
