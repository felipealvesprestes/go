package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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
			exibeLogs()
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

	websites := leArquivoListaWebsites()

	for i := 0; i < quantidadeMonitoramento; i++ {
		for _, website := range websites {
			testaWebsite(website)
		}

		time.Sleep(delay * time.Second)

		fmt.Println()
	}
}

func testaWebsite(website string) {
	status, erro := http.Get(website)

	trataMensagemErro(erro)

	if status.StatusCode == 200 {
		fmt.Println("[", website, "] está online.")
		gravaLog(website, true)
	} else {
		gravaLog(website, false)
		fmt.Println("[", website, "] está fora do ar.")
	}
}

func leArquivoListaWebsites() []string {
	var websites []string

	arquivo, erro := os.Open("lista-websites.txt")

	trataMensagemErro(erro)

	buffer := bufio.NewReader(arquivo)

	for {
		linha, erro := buffer.ReadString('\n')

		linha = strings.TrimSpace(linha)

		websites = append(websites, linha)

		if erro == io.EOF {
			break
		}
	}

	arquivo.Close()

	return websites
}

func trataMensagemErro(erro error) {
	if erro != nil {
		fmt.Println("Erro#", erro)
		os.Exit(-1)
	}
}

func gravaLog(website string, status bool) {
	arquivo, erro := os.OpenFile("log.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	trataMensagemErro(erro)

	arquivo.WriteString("[" + time.Now().Format("02/01/2006 15:04:05") + "] " + website + " (online: " + strconv.FormatBool(status) + ")\n")
	arquivo.Close()
}

func exibeLogs() {
	arquivo, erro := ioutil.ReadFile("log.txt")

	trataMensagemErro(erro)

	fmt.Println(string(arquivo))
}
