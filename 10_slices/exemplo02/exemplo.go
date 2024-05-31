package main

import "fmt"

func main() {
	impares := [7]int{1, 3, 5, 7, 9, 11, 13}
	fmt.Println(impares)

	fatia := impares[2:5]
	fmt.Println(fatia)

	fatia[0] = 21
	fmt.Println(fatia)

	fmt.Println(impares)
}
