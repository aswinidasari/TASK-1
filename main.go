package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	counter int64
	mu      sync.Mutex
)

func counterHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	counter++
	fmt.Fprintf(w, "%d", counter)
	mu.Unlock()
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().UTC()
	fmt.Fprintf(w, currentTime.Format(time.RFC3339))
}

func main() {
	http.HandleFunc("/v1/counter", counterHandler)
	http.HandleFunc("/v1/time", timeHandler)

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
