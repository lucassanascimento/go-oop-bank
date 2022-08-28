package contas

import c "github.com/go-bank/clientes"

type ContaPoupanca struct {
	Titular       c.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (c *ContaPoupanca) Sacar(valor float64) string {
	podeSacar := valor <= c.saldo && valor > 0
	if podeSacar {
		c.saldo -= valor
		return "Sucess"
	} else {
		return "Failure"
	}

}

func (c *ContaPoupanca) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "Depósido Sucesso :)", c.saldo
	} else {
		return "Depósido falha :<", c.saldo
	}
}

func (c *ContaPoupanca) Saldo() float64 {
	return c.saldo
}
