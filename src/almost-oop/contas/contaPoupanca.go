package contas

import "almost-oop/clientes"

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (conta *ContaPoupanca) Sacar(valorSaque float64) (string, float64) {

	podeSacar := valorSaque > 0 && valorSaque <= conta.saldo

	if podeSacar {
		conta.saldo -= valorSaque
		return "Saque efetuado.", conta.saldo
	}

	return "Saldo insuficiente.", conta.saldo
}

func (conta *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito > 0 {
		conta.saldo += valorDeposito
		return "Deposito efetuado.", conta.saldo
	}

	return "Valor do depósito menor que 0.", conta.saldo
}

func (conta *ContaPoupanca) GetSaldo() float64 {
	return conta.saldo
}
