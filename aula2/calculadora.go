package main

import "fmt"

func main() {

	/*
		var (
			num1     int
			num2     int
			operador string

		)
	*/

	for i := 0; i <= 10; i += 2 {
		fmt.Println(i)
	}

	contador := 0
	for {
		fmt.Println(contador, "Este é um loop infinito")
		contador++
		if contador >= 5 {
			break
		}
	}
	fmt.Println(contador, "Fim do loop")

	//fmt.Println("Bem vinda a calculadora")

	pontuacoes := []int{99, 4, -5, 22, -16}
	var total int
	for indice, p := range pontuacoes {
		if p < 0 {
			fmt.Println(indice, ": pontuação negativa desconsiderada")
			continue
		}
		total += p
	}

}
