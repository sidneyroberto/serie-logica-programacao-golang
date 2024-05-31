package main

import "fmt"

func main() {
	nomes := [5]string{"Cleuza", "Lisbela", "Elizabete", "Artur", "Leonardo"}

	mulheres := nomes[0:3]
	fmt.Println(mulheres)

	homens := nomes[3:5]
	fmt.Println(homens)

	mulheres = nomes[:3]
	fmt.Println(mulheres)

	homens = nomes[3:]
	fmt.Println(homens)
}
