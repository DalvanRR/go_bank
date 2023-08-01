package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (conta *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= conta.saldo
	if podeSacar {
		conta.saldo -= valorDoSaque
		return "Saque realizado com sucesso", conta.saldo
	} else {
		return "Saldo insuficiente", conta.saldo
	}
}

func (conta *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		conta.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", conta.saldo
	} else {
		return "Valor do depósito menor que zero", conta.saldo
	}
}

func main() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia.saldo)
	fmt.Println(contaDaSilvia.Depositar(2000))
}
