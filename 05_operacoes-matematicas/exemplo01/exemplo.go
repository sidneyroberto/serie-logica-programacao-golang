package main

import "fmt"

func main() {
	fmt.Println("Digite o primeiro número:")
	var numero1 int
	fmt.Scanf("%d\n", &numero1)

	fmt.Println("Digite o segundo número:")
	var numero2 int
	fmt.Scanf("%d\n", &numero2)

	soma := numero1 + numero2
	subtracao := numero1 - numero2
	multiplicacao := numero1 * numero2
	divisao := numero1 / numero2
	modulo := numero1 % numero2

	fmt.Printf("%d + %d = %d\n", numero1, numero2, soma)
	fmt.Printf("%d - %d = %d\n", numero1, numero2, subtracao)
	fmt.Printf("%d * %d = %d\n", numero1, numero2, multiplicacao)
	fmt.Printf("%d / %d = %d\n", numero1, numero2, divisao)
	fmt.Printf("%d %% %d = %d\n", numero1, numero2, modulo)
}
