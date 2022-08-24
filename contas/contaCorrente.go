package contas

import c "github.com/go-bank/clientes"

type ContaCorrente struct {
	Titular  c.Titular
	NAgencia int
	Conta    int
	saldo    float64
}

func Somando(numeros ...int) int {
	resultadoDaSoma := 0
	for _, numero := range numeros {
		resultadoDaSoma += numero
	}
	return resultadoDaSoma
}

func (c *ContaCorrente) Sacar(valor float64) string {
	podeSacar := valor <= c.saldo && valor > 0
	if podeSacar {
		c.saldo -= valor
		return "Sucess"
	} else {
		return "Failure"
	}

}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "Depósido Sucesso :)", c.saldo
	} else {
		return "Depósido falha :<", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valor float64, to *ContaCorrente) string {
	if valor <= c.saldo && valor > 0 {
		c.saldo -= valor
		to.Depositar(valor)
		return "Sucesso ao transferir"
	} else {
		return "Erro ao transferir"
	}
}

func (c *ContaCorrente) Saldo() float64 {
	return c.saldo
}
