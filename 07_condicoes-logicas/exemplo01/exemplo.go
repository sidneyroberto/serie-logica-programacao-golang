package main

import "fmt"

func main() {
	fmt.Println("Digite o dividendo:")
	var dividendo float64
	fmt.Scanf("%f\n", &dividendo)

	fmt.Println("Digite o divisor:")
	var divisor float64
	fmt.Scanf("%f\n", &divisor)

	if divisor != 0 {
		resultado := dividendo / divisor
		fmt.Printf("%f / %f = %f\n", dividendo, divisor, resultado)
	} else {
		fmt.Println("Divisor não pode ser igual a 0")
	}
}
