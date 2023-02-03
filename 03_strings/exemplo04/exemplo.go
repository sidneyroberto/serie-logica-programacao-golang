package main

import (
	"fmt"
	"strings"
)

func main() {
	frase := "Um delicioso bolo de laranja"
	fmt.Printf("Frase original: %s\n", frase)
	fraseMinuscula := strings.ToLower(frase)
	fmt.Printf("Frase em minúsculo: %s\n", fraseMinuscula)
	fraseMaiuscula := strings.ToUpper(frase)
	fmt.Printf("Frase em maiúsculo: %s\n", fraseMaiuscula)

	outraFrase := "   uma frase no meio   "
	tamanho := len(outraFrase)
	fmt.Printf("Tamanho da frase: %d\n", tamanho)
	outraFrase = strings.Trim(outraFrase, " ")
	tamanho = len(outraFrase)
	fmt.Printf("Tamanho da frase depois: %d\n", tamanho)
}
