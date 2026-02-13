package main

import (
	"errors"
	"fmt"
)

type Transaction struct {
	ID     string
	Amount float64
}
type Ledger struct {
	SeenIDs map[string]bool
}

func (l *Ledger) ProcessTrans(t Transaction) (string, error) {
	if l.SeenIDs[t.ID] {
		return "", errors.New("duplicate transaction ID")
	}
	if t.Amount <= 0 {
		return "", errors.New("Invalid amount, Retry.")
	}
	l.SeenIDs[t.ID] = true
	return fmt.Sprintf("Transaction Processed: $%.2f", t.Amount), nil
}

func main() {
	myLedger := Ledger{SeenIDs: make(map[string]bool)}
	t1 := Transaction{ID: "Robot1", Amount: 500.00}
	t2 := Transaction{ID: "Robot1", Amount: 500.00}

	msg, err := myLedger.ProcessTrans(t1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(msg)
	}
	msg2, err2 := myLedger.ProcessTrans(t2)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println(msg2)
	}
}
