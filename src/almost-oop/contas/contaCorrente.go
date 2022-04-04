package contas

import "almost-oop/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (conta *ContaCorrente) Sacar(valorSaque float64) (string, float64) {

	podeSacar := valorSaque > 0 && valorSaque <= conta.Saldo

	if podeSacar {
		conta.Saldo -= valorSaque
		return "Saque efetuado.", conta.Saldo
	}

	return "Saldo insuficiente.", conta.Saldo
}

func (conta *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito > 0 {
		conta.Saldo += valorDeposito
		return "Deposito efetuado.", conta.Saldo
	}

	return "Valor do depósito menor que 0.", conta.Saldo
}

func (conta *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {

	if valorTransferencia < conta.Saldo && valorTransferencia > 0 {
		conta.Saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	}

	return false
}
