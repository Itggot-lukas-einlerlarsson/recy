package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"
)

// RecycleOrder represents a pickup request
type RecycleOrder struct {
	Name    string
	Address string
	Items   string // e.g., "Plastic, Paper"
	Date    string // e.g., "2025-10-10"
}

var (
	templates   *template.Template
	orders      []RecycleOrder
	ordersMutex sync.Mutex
)

func main() {
	// Pre-compile templates
	var err error
	templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Failed to load templates: %v", err)
	}

	// Routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/order", formHandler)
	http.HandleFunc("/submit", submitHandler)
	http.HandleFunc("/queue", queueHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handlers (updated to use pre-compiled templates)
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "form.html", nil); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func queueHandler(w http.ResponseWriter, r *http.Request) {
	ordersMutex.Lock()
	defer ordersMutex.Unlock()

	if err := templates.ExecuteTemplate(w, "queue.html", orders); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Create order
	order := RecycleOrder{
		Name:    r.FormValue("name"),
		Address: r.FormValue("address"),
		Items:   r.FormValue("items"),
		Date:    r.FormValue("date"),
	}

	// Add to queue
	ordersMutex.Lock()
	orders = append(orders, order)
	ordersMutex.Unlock()

	// Redirect to confirmation
	http.Redirect(w, r, "/queue?status=submitted", http.StatusSeeOther)
}
