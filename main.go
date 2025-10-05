package main

import (
	"fmt"
	"net/http"
	"sync"
)

type RecycleOrder struct {
	Name    string
	Address string
	Items   string
	Date    string
}

var (
	orders      []RecycleOrder
	ordersMutex sync.Mutex
)

func main() {
	// Define a route handler for the root path
	http.HandleFunc("/recycle", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/index.html")
	})

	// Start the server
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
