package main

import "fmt"

func main() {

	var num1, num2 int
	num1 = 10

	//Perguntar numero pro user
	fmt.Println("Hey, me diz um numero aí?")

	//pega o numero do user
	fmt.Scan(&num2)

	// somar os numeros. := é atribuição e é uma forma de inicializar
	result := num1 + num2
	fmt.Println(result)

}
