# Go Ledger (Concurrent Ingestion)
### *Bilingual Logic & Infrastructure Study*

## Overview
This project is a high-performance **Concurrent Ledger** built in **Go (Golang)**. It serves as a "Language Agnostic" demonstration of the **Idempotency** and **Deduplication** logic originally developed in C#, now optimized for Go's lightweight concurrency model.

## Key Technical Features
- **Goroutine Workers:** Implements a background consumer using Go's lightweight threading model for non-blocking transaction processing.
- **Channel-Based Communication:** Utilizes **Buffered Channels** to facilitate thread-safe memory sharing between producers and consumers (CSP Pattern).
- **Explicit Error Handling:** Demonstrates Go's "Comma-Ok" idiom for robust, traceable error management without exception bubbling.
- **Idempotency Mapping:** Employs Go's native `map` types to ensure $O(1)$ deduplication performance.

## Why Go?
To demonstrate true **Language Agnosticism**, I ported my core financial logic to Go to take advantage of its low-level efficiency and unique approach to concurrency (CSP). This proves that architectural principles (Idempotency, Resiliency, Persistence) transcend specific syntax.
