package main

import (
	"fmt"
	"math"
)

func main() {
	numeros := [4]float64{5, 4, 9, 3}

	for i, v := range numeros {
		resultado := math.Pow(numeros[i], 2)
		fmt.Printf("%f^2 = %f\n", v, resultado)
	}
}
