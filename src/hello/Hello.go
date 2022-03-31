package main

import (
	"fmt"
)

func main() {

	nome := "Felipe"
	versao := 1.1

	fmt.Println("Olá", nome)
	fmt.Println("A versão do programa é", versao)

	fmt.Println("1)  Inicicar monitoramento")
	fmt.Println("2)  Exibir logs")
	fmt.Println("3)  Encerrar programa")

	var opcao int
	fmt.Scan(&opcao)
	fmt.Println("O comando escolhido foi", opcao)

	switch opcao {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 3:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Por favor, selecione uma opção válida.")
	}
}
