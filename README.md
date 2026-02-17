# Go Ledger (Concurrent SQL Ingestion)
### *Bilingual Logic & High-Performance Infrastructure Study*

## Overview
A high-performance **Concurrent Ledger** built in **Go (Golang)**. This project demonstrates the porting of core **Idempotency** and **Deduplication** logic from C# into a low-level, memory-efficient environment. It serves as a proof of **Language Agnosticism** and architectural adaptability.

## Key Technical Features
- **Goroutine Worker Pool:** Utilizes Go's lightweight threading model to process transactions in the background with minimal overhead.
- **CSP Pattern (Channels):** Employs **Buffered Channels** to facilitate thread-safe communication between the producer and the SQL consumer.
- **Relational Persistence (Pure Go):** Integrated with a **Pure-Go SQLite driver** to ensure zero-dependency portability across Windows, Linux, and macOS.
- **Native HTTP API:** Implements a built-in status monitoring endpoint using the Go `net/http` standard library, demonstrating a "Zero-Framework" approach to Web Services.
- **Atomic Deduplication:** Leverages **SQL Primary Key constraints** to enforce data integrity at the database level.

## Why Go?
By rebuilding the transaction logic in Go, I’ve demonstrated that architectural principles like **Persistence**, **Concurrency**, and **Fault Tolerance** transcend specific syntax. This implementation focuses on Go's unique approach to "Sharing memory by communicating," providing a stark contrast to the Object-Oriented patterns used in my .NET projects.

## How to Run
1. Ensure [Go 1.21+](https://go.dev) is installed.
2. `go run main.go`
3. Navigate to `http://localhost:8080/status` to view real-time ledger metrics.
