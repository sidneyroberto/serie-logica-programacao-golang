package main

import "fmt"

func main() {
	var valor1 bool
	valor1 = true
	valor2 := true

	resultado := valor1 && valor2
	fmt.Printf("%t && %t = %t", valor1, valor2, resultado)

	resultadoAnd1 := true && true
	resultadoAnd2 := true && false
	resultadoAnd3 := false && true
	resultadoAnd4 := false && false
	fmt.Printf("\n\ntrue && true = %t\n", resultadoAnd1)
	fmt.Printf("true && false = %t\n", resultadoAnd2)
	fmt.Printf("false && true = %t\n", resultadoAnd3)
	fmt.Printf("false && false = %t\n", resultadoAnd4)

	resultado = valor1 || valor2
	fmt.Printf("\n%t || %t = %t", valor1, valor2, resultado)

	resultadoOr1 := true || true
	resultadoOr2 := true || false
	resultadoOr3 := false || true
	resultadoOr4 := false || false

	fmt.Printf("\n\ntrue || true = %t\n", resultadoOr1)
	fmt.Printf("true || false = %t\n", resultadoOr2)
	fmt.Printf("false || true = %t\n", resultadoOr3)
	fmt.Printf("false || false = %t\n", resultadoOr4)

	negacaoValor1 := !valor1
	fmt.Printf("\n\nNegação de %t: %t", valor1, negacaoValor1)
}
