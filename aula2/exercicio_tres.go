package main

import "fmt"

func main() {

	var num3 int

	fmt.Println("Manda um número aí")

	fmt.Scan(&num3)

	if num3 > 0 {
		fmt.Println("Positivo")

	} else if num3 < 0 {
		fmt.Println("Negativo")

	} else {
		fmt.Println("Zero!")
	}

}
