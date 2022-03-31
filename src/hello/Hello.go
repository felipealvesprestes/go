package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()
	exibeMenu()

	opcao := leOpcao()

	switch opcao {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 3:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Opção válida.")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	nome := "Felipe"
	versao := 1.1

	fmt.Println("Olá", nome)
	fmt.Println("A versão do programa é", versao)
}

func exibeMenu() {
	fmt.Println("1)  Inicicar monitoramento")
	fmt.Println("2)  Exibir logs")
	fmt.Println("3)  Encerrar programa")
}

func leOpcao() int {
	var opcao int
	fmt.Scan(&opcao)
	fmt.Println("O comando escolhido foi", opcao)

	return opcao
}
