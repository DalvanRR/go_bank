package main

import (
	"fmt"

	"github.com/go_bank/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDalvan := contas.ContaPoupanca{}
	contaDoDalvan.Depositar(100)
	PagarBoleto(&contaDoDalvan, 60)

	fmt.Println(contaDoDalvan.ObterSaldo())

	contaDoOdair := contas.ContaCorrente{}
	contaDoOdair.Depositar(500)
	PagarBoleto(&contaDoOdair, 1000)

	fmt.Println(contaDoOdair.ObterSaldo())

}
