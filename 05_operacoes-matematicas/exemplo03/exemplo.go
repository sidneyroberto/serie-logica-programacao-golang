package main

import "fmt"

func main() {
	fmt.Println("Digite o valor a ser emprestado:")
	var valorEmprestimo float64
	fmt.Scanf("%f\n", &valorEmprestimo)

	fmt.Println("Digite o valor da taxa de juros:")
	var taxaJuros float64
	fmt.Scanf("%f\n", &taxaJuros)

	fmt.Println("Digite a quantidade de parcelas:")
	var quantidadeParcelas int
	fmt.Scanf("%d\n", &quantidadeParcelas)

	montante := valorEmprestimo * (1 + ((taxaJuros / 100) * float64(quantidadeParcelas)))
	valorPrestacao := montante / float64(quantidadeParcelas)

	fmt.Printf("Valor total a ser pago: %.2f\n", montante)
	fmt.Printf("Valor de cada prestação: %.2f\n", valorPrestacao)

}
