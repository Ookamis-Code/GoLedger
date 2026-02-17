package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS ledger (id TEXT PRIMARY KEY, amount REAL)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM ledger").Scan(&count)
		if err != nil {
			http.Error(w, "Failed to retrieve status", http.StatusInternalServerError)
			return
		}
		repsonse := map[string]interface{}{
			"status":             "Active",
			"total_transactions": count,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(repsonse)
	})
	fmt.Println("Server is running on http://localhost:8080/status")
	go http.ListenAndServe(":8080", nil)
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()
	queue := make(chan Transaction)
	go func() {
		for ts := range queue {
			stmt, err := db.Prepare("INSERT INTO ledger (id, amount) VALUES (?, ?)")
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
	queue <- Transaction{ID: "Robit2", Amount: 150.0}
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
