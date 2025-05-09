# Golang: Concurrent HTML & Text Processing with sync.Mutex and sync.RWMutex

This is an educational project built with Go to explore and demonstrate safe concurrent programming using `sync.Mutex` and `sync.RWMutex`. The project consists of two independent modules:

- [**HTML Parser with RWMutex**](internal/RWMutex/HtmlParser.go): Parses HTML pages concurrently and extracts content by specified tags (like `<h1>`, `<script>`, etc.), while managing shared memory using `RWMutex`.
- [**Letter Statistics with Mutex**](internal/mutex_tr/LetterStat_Parallel.go): Calculates character frequencies in text files using parallel processing with `Mutex` for safe shared access.

---
## Project Goals

- Learn and practice safe concurrent programming in Go.
- Demonstrate the use of low-level synchronization primitives.
- Write and manage concurrent parsing logic by hand without relying on third-party libraries.
- Understand how race conditions and shared memory access can be controlled using the Go standard library.

---
## About Manual Tag Parsing
1. Why did I parse HTML manually instead of using a library like golang.org/x/net/html or goquery? 
   - This was a conscious decision for learning purposes.I wanted to better understand how tag boundaries work in raw HTML.

### This also triggered interesting scenarios such as:
1. Antiviruses (like Kaspersky) falsely flagging the binary when manual parsing logic looked suspicious (Trojan).
2. Race condition exposure when multiple goroutines read and write shared data structures without locks.
