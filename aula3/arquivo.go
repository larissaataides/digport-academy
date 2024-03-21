package main

//import de libs/ bibliotecas
import "fmt"

var variavelGlobal string

func FunçãoPublica() {
	variavel := "SóDessaFunçãoPublica"
	funçãoPrivada()
	fmt.Print(variavel)

	variavelGlobal = "variavelNaPublica"
}

func funçãoPrivada() {
	variavel := "SóDessaFunçãoPrivada"
	fmt.Print(variavel)
	//Escolpo da função
	variavelGlobal = "variavelNaPrivada"
}
