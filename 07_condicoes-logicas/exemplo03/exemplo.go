package main

import "fmt"

func main() {
	fmt.Println("Digite o número 1:")
	var numero1 float64
	fmt.Scanf("%f\n", &numero1)

	fmt.Println("Digite o número 2:")
	var numero2 float64
	fmt.Scanf("%f\n", &numero2)

	fmt.Println("Digite a opção desejada:")
	fmt.Println("1 - Soma")
	fmt.Println("2 - Subtração")
	fmt.Println("3 - Multiplicação")
	fmt.Println("4 - Divisão")
	var opcao string
	fmt.Scanf("%s", &opcao)

	var resultado float64
	resultadoFoiCalculado := true

	switch opcao {
	case "1":
		resultado = numero1 + numero2
	case "2":
		resultado = numero1 - numero2
	case "3":
		resultado = numero1 * numero2
	case "4":
		if numero2 != 0 {
			resultado = numero1 / numero2
		} else {
			fmt.Println("Divisor não pode ser igual a 0")
			resultadoFoiCalculado = false
		}
	default:
		fmt.Println("Opção inválida")
		resultadoFoiCalculado = false
	}

	if resultadoFoiCalculado {
		fmt.Printf("Resultado: %f\n", resultado)
	}
}
