package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoDalvan := ContaCorrente{titular: "Dalvan", numeroAgencia: 589, numeroConta: 123456, saldo: 125.50}
	contaDoOdair := ContaCorrente{"Odair", 222, 654321, 125.50}

	fmt.Println(contaDoDalvan)
	fmt.Println(contaDoOdair)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	fmt.Println(*contaDaCris)
}
