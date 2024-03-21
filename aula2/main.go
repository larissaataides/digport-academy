package main

import (
	"fmt"
	"os"
)

func main() {

	var name string

	//Nome do usuário
	fmt.Println("Hey, qual o seu nome?")

	//pega o nome para a variavel
	fmt.Scan(&name)

	//imprime na tela o nome digitado
	fmt.Fprintln(os.Stdout, []any{"Bem vindo " + name}...)
	//fmt.Printf("Bem vindo, ",)
}
