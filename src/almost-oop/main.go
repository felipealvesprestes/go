package main

import (
	"almost-oop/clientes"
	"almost-oop/contas"
	"fmt"
)

func main() {

	cliente1 := clientes.Titular{}
	cliente1.Nome = "Felipe Prestes"
	cliente1.CPF = "022.890.778-45"
	cliente1.Profissao = "Desenvolvedor"

	cliente2 := clientes.Titular{}
	cliente2.Nome = "Pedro Silva"
	cliente2.CPF = "352.120.732-35"
	cliente2.Profissao = "Advogado"

	cliente3 := clientes.Titular{}
	cliente3.Nome = "Amanda Rodrigues"
	cliente3.CPF = "876.355.233-67"
	cliente3.Profissao = "Farmacêutica"

	contaFelipe := contas.ContaCorrente{}

	contaFelipe.NumeroAgencia = 598
	contaFelipe.NumeroConta = 3322
	contaFelipe.Titular = cliente1
	contaFelipe.Depositar(3450)

	contaPedro := contas.ContaCorrente{}

	contaPedro.NumeroAgencia = 778
	contaPedro.NumeroConta = 4677
	contaPedro.Titular = cliente2
	contaPedro.Depositar(3000)

	contaAmanda := contas.ContaPoupanca{}

	contaAmanda.NumeroAgencia = 233
	contaAmanda.NumeroConta = 8799
	contaAmanda.Titular = cliente3
	contaAmanda.Depositar(2000)

	fmt.Println("Saldo atual:", contaFelipe.GetSaldo())

	retornoSaque, saldoAposSaque := contaFelipe.Sacar(100)
	fmt.Println(retornoSaque, "Saldo atual:", saldoAposSaque)

	retornoDeposito, saldoAposDeposito := contaFelipe.Depositar(1000)
	fmt.Println(retornoDeposito, "Saldo atual:", saldoAposDeposito)

	if contaFelipe.Transferir(350, &contaPedro) {
		fmt.Println("Transferencia realizada.")
		fmt.Println("Saldo atual (Felipe):", contaFelipe.GetSaldo())
		fmt.Println("Saldo atual (Pedro):", contaPedro.GetSaldo())
	} else {
		fmt.Println("Saldo insuficiente.")
	}

	retornoDepositoPoupanca, _ := contaAmanda.Depositar(400)

	contaAmanda.Sacar(50)

	fmt.Println(retornoDepositoPoupanca, "Saldo atual:", contaAmanda.GetSaldo())

	PagarBoleto(&contaAmanda, 100)
	fmt.Println("Boleto pago. Saldo atual (Amanda):", contaAmanda.GetSaldo())

	PagarBoleto(&contaFelipe, 100)
	fmt.Println("Boleto pago. Saldo atual (Felipe):", contaFelipe.GetSaldo())
}

func PagarBoleto(conta VerificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type VerificarConta interface {
	Sacar(valorSaque float64) (string, float64)
}
