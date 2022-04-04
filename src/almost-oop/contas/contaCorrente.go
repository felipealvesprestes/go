package contas

import "almost-oop/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (conta *ContaCorrente) Sacar(valorSaque float64) (string, float64) {

	podeSacar := valorSaque > 0 && valorSaque <= conta.saldo

	if podeSacar {
		conta.saldo -= valorSaque
		return "Saque efetuado.", conta.saldo
	}

	return "Saldo insuficiente.", conta.saldo
}

func (conta *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {

	if valorDeposito > 0 {
		conta.saldo += valorDeposito
		return "Deposito efetuado.", conta.saldo
	}

	return "Valor do depósito menor que 0.", conta.saldo
}

func (conta *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {

	if valorTransferencia < conta.saldo && valorTransferencia > 0 {
		conta.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	}

	return false
}

func (conta *ContaCorrente) GetSaldo() float64 {
	return conta.saldo
}
