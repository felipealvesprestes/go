package main

import (
	"fmt"
	"reflect"
)

func main() {

	nome := "Felipe"
	sobrenome := "Prestes"
	idade := 27
	versao := 1.1

	fmt.Println("Olá", nome, sobrenome)
	fmt.Println("Sua idade é", idade)
	fmt.Println("A versão do programa é", versao)

	fmt.Println("O tipo da variável sobrenome é", reflect.TypeOf(sobrenome))
}
