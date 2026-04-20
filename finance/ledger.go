package finance

import "errors"

var (
	ErrInvalidAmount          = errors.New("amount must be greater than zero")
	ErrInvalidTransactionType = errors.New("transaction type must be income or expense")
)

type TransactionType string

const (
	Income  TransactionType = "income"
	Expense TransactionType = "expense"
)

type Transaction struct {
	Description string
	Amount      float64
	Type        TransactionType
}

type Ledger struct {
	transactions []Transaction
}

func NewLedger() *Ledger {
	return &Ledger{transactions: make([]Transaction, 0)}
}

func (l *Ledger) AddTransaction(transaction Transaction) error {
	if transaction.Amount <= 0 {
		return ErrInvalidAmount
	}

	if transaction.Type != Income && transaction.Type != Expense {
		return ErrInvalidTransactionType
	}

	l.transactions = append(l.transactions, transaction)
	return nil
}

func (l *Ledger) Balance() float64 {
	var balance float64
	for _, transaction := range l.transactions {
		if transaction.Type == Income {
			balance += transaction.Amount
			continue
		}

		balance -= transaction.Amount
	}

	return balance
}

func (l *Ledger) TotalByType(transactionType TransactionType) float64 {
	var total float64
	for _, transaction := range l.transactions {
		if transaction.Type == transactionType {
			total += transaction.Amount
		}
	}

	return total
}

func (l *Ledger) Transactions() []Transaction {
	transactions := make([]Transaction, len(l.transactions))
	copy(transactions, l.transactions)
	return transactions
}
