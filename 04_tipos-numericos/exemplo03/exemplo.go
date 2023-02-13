package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Digite a sua altura:")
	var altura float64
	fmt.Scanf("%f\n", &altura)

	fmt.Println("Digite o seu peso:")
	var peso float64
	fmt.Scanf("%f\n", &peso)

	imc := peso / math.Pow(altura, 2)
	fmt.Printf("O seu IMC: %f", imc)
}
