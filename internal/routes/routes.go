package routes

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"trash-pickup-app/internal/handlers"
)


func servePickupForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "static/pickup_form.html")
}

func SetupRoutes(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	handler := &handlers.PickupHandler{DB: db}

	r.Get("/", servePickupForm)
	r.Post("/pickup", handler.CreatePickup)
	r.Get("/pickup_list", handler.GetAllPickups)

	return r
}

