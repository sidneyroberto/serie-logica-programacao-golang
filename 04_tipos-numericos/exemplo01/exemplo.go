package main

import "fmt"

func main() {
	// uint32 -> 0 a 4.294.967.295
	// uint64 -> 0 a 18.446.744.073.709.551.615
	// int32 -> -2.147.483.648 a 2.147.483.647
	// int64 -> -9.223.372.036.854.775.808 a 9.223.372.036.854.775.807
	var quantidade uint = 34
	var temperatura int = -3
	fmt.Printf("Quantidade: %d\nTemperatura: %d\n", quantidade, temperatura)

	//float32	->	-3.4e+38 to 3.4e+38.
	//float64	->	-1.7e+308 to +1.7e+308.
	var peso float64 = 1.2
	fmt.Printf("Peso: %.0f\n", peso)
}
