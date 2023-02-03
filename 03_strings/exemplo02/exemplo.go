package main

import "fmt"

func main() {
	frase := "Olá, mundo!"
	primeiroCaractere := frase[0]
	fmt.Printf("Primeiro caractere: %c\n", primeiroCaractere)
	indiceUltimoCaractere := len(frase) - 1
	ultimoCaractere := frase[indiceUltimoCaractere]
	fmt.Printf("Último caractere: %c\n", ultimoCaractere)
}
