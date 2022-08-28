package main

import (
	"fmt"

	"github.com/go-bank/contas"
)

func PagarBoleto(conta varificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type varificarConta interface {
	Sacar(valor float64) string
}

func main() {
	lucas := contas.ContaCorrente{}

	lucas.Depositar(10)
	PagarBoleto(&lucas, 6)
	fmt.Println(lucas.Saldo())
}
