package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaMaria := ContaCorrente{
		"Maria Josefina",
		598,
		3322,
		4578.30,
	}

	contaJoao := ContaCorrente{
		"João Silva",
		598,
		3322,
		4578.30,
	}

	fmt.Println(contaMaria.titular)
	fmt.Println(contaJoao.titular)
}
