package main

import (
	"fmt"

	"github.com/go-bank/contas"
)

func main() {
	// clientLucas := clientes.Titular{Nome: "Lucas", CPF: "03849593", Profissao: "dev"}
	lucas := contas.ContaCorrente{}

	fmt.Println(lucas.Depositar(10))
	fmt.Println(lucas.Saldo())
}
