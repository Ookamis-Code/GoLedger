package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/glebarez/sqlite"
)

type Transaction struct {
	ID     string
	Amount float64
}

func main() {
	db, err := sql.Open("sqlite", "./ledger.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS transactions (id TEXT PRIMARY KEY, amount REAL)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err)
	}
	queue := make(chan Transaction)

	go func() {
		for ts := range queue {
			stmt, err := db.Prepare("INSERT INTO transactions (id, amount) VALUES (?, ?)")
			if err != nil {
				fmt.Printf("Failed to prepare statement: %v\n", err)
				continue
			}
			_, err = stmt.Exec(ts.ID, ts.Amount)
			if err != nil {
				fmt.Printf("Failed to insert transaction %s: %v\n", ts.ID, err)
			} else {
				fmt.Printf("Inserted transaction %s successfully\n", ts.ID)
			}
		}
	}()
	queue <- Transaction{ID: "Robit1", Amount: 100.0}
	queue <- Transaction{ID: "Robit1", Amount: 100.0}
	close(queue)
	fmt.Scanln()
}
