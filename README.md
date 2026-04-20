# go-finance

Projeto financeiro em Golang com um ledger simples para registrar receitas e despesas.

## Funcionalidades

- Registro de transações de receita (`income`) e despesa (`expense`)
- Cálculo de saldo
- Cálculo de total por tipo de transação

## Exemplo de uso

```go
package main

import (
	"fmt"

	"github.com/blackbarth/go-finance/finance"
)

func main() {
	ledger := finance.NewLedger()
	_ = ledger.AddTransaction(finance.Transaction{Description: "Salário", Amount: 5000, Type: finance.Income})
	_ = ledger.AddTransaction(finance.Transaction{Description: "Aluguel", Amount: 1500, Type: finance.Expense})

	fmt.Println("Saldo:", ledger.Balance())
}
```

## Testes

```bash
go test ./...
```
