package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	ID     string
	Amount float64
}
type Ledger struct {
	SeenIDs map[string]bool
}

func main() {
	queue := make(chan Transaction)
	go func() {
		seen := make(map[string]bool)
		for ts := range queue {
			if seen[ts.ID] {
				fmt.Printf("Duplicate transaction ID: %s\n", ts.ID)
			} else {
				seen[ts.ID] = true
				fmt.Printf("Processed transaction: %s\n", ts.ID)
			}
		}
	}()
	queue <- Transaction{ID: "Robot1", Amount: 100.0}
	queue <- Transaction{ID: "Robot2", Amount: 200.0}
	queue <- Transaction{ID: "Robot1", Amount: 100.0}
	time.Sleep(500 * time.Millisecond)
	close(queue)
}
