package contas

import "github.com/go_bank/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (conta *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= conta.saldo
	if podeSacar {
		conta.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (conta *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		conta.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", conta.saldo
	} else {
		return "Valor do deposito menor que zero", conta.saldo
	}
}

func (conta *ContaPoupanca) ObterSaldo() float64 {
	return conta.saldo
}
