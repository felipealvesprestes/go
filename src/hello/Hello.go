package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const quantidadeMonitoramento = 2
const delay = 3

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

		opcao := leOpcao()

		switch opcao {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 3:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Encerrando programa.")
			fmt.Println()
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Felipe"
	versao := 1.1

	fmt.Println("+-------------------------")
	fmt.Println("| Olá", nome)
	fmt.Println("|")
	fmt.Println("| Versão:", versao)
	fmt.Println("+-------------------------")
	fmt.Println()
}

func exibeMenu() {
	fmt.Println("1)  Iniciar monitoramento")
	fmt.Println("2)  Exibir logs")
	fmt.Println("3)  Encerrar programa")
	fmt.Println()
}

func leOpcao() int {
	var opcao int
	fmt.Scan(&opcao)
	fmt.Println()
	fmt.Println("Opção selecionada: ", opcao)
	fmt.Println()

	return opcao
}

func iniciarMonitoramento() {

	websites := []string{
		"http://eelslap.com",
		"https://random-status-code.herokuapp.com",
		"https://google.com.br",
	}

	for i := 0; i < quantidadeMonitoramento; i++ {
		for _, website := range websites {
			testaWebsite(website)
		}

		time.Sleep(delay * time.Second)

		fmt.Println()
	}
}

func testaWebsite(website string) {
	status, _ := http.Get(website)

	if status.StatusCode == 200 {
		fmt.Println("[", website, "] está online.")
	} else {
		fmt.Println("[", website, "] está fora do ar.")
	}
}
