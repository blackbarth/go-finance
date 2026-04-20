package finance

import "testing"

func TestLedgerAddTransactionAndBalance(t *testing.T) {
	ledger := NewLedger()

	if err := ledger.AddTransaction(Transaction{Description: "salary", Amount: 3000, Type: Income}); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if err := ledger.AddTransaction(Transaction{Description: "rent", Amount: 1200, Type: Expense}); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if got, want := ledger.TotalByType(Income), 3000.0; got != want {
		t.Fatalf("income total mismatch: got %v, want %v", got, want)
	}

	if got, want := ledger.TotalByType(Expense), 1200.0; got != want {
		t.Fatalf("expense total mismatch: got %v, want %v", got, want)
	}

	if got, want := ledger.Balance(), 1800.0; got != want {
		t.Fatalf("balance mismatch: got %v, want %v", got, want)
	}
}

func TestLedgerAddTransactionValidation(t *testing.T) {
	ledger := NewLedger()

	if err := ledger.AddTransaction(Transaction{Description: "invalid amount", Amount: 0, Type: Income}); err != ErrInvalidAmount {
		t.Fatalf("expected ErrInvalidAmount, got %v", err)
	}

	if err := ledger.AddTransaction(Transaction{Description: "invalid type", Amount: 100, Type: "other"}); err != ErrInvalidTransactionType {
		t.Fatalf("expected ErrInvalidTransactionType, got %v", err)
	}
}

func TestTransactionsReturnsCopy(t *testing.T) {
	ledger := NewLedger()
	if err := ledger.AddTransaction(Transaction{Description: "salary", Amount: 1000, Type: Income}); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	transactions := ledger.Transactions()
	transactions[0].Amount = 9999

	if got, want := ledger.Transactions()[0].Amount, 1000.0; got != want {
		t.Fatalf("expected ledger state to remain immutable via accessor, got %v, want %v", got, want)
	}
}
