package main

import "fmt"

func main() {
	numero1 := 7
	numero2 := 10

	resultado1 := numero1 == numero2
	resultado2 := numero1 > numero2
	resultado3 := numero1 < numero2
	resultado4 := numero1 >= numero2
	resultado5 := numero1 <= numero2
	resultado6 := numero1 != numero2

	fmt.Printf("%d == %d = %t\n", numero1, numero2, resultado1)
	fmt.Printf("%d > %d = %t\n", numero1, numero2, resultado2)
	fmt.Printf("%d < %d = %t\n", numero1, numero2, resultado3)
	fmt.Printf("%d >= %d = %t\n", numero1, numero2, resultado4)
	fmt.Printf("%d <= %d = %t\n", numero1, numero2, resultado5)
	fmt.Printf("%d != %d = %t\n", numero1, numero2, resultado6)

	resultado := (98 < 99) && (true && false) || true
	fmt.Printf("\nResultado: %t\n", resultado)
}
