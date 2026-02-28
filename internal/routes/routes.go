package routes

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"trash-pickup-app/internal/handlers"
)

func SetupRoutes(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	handler := &handlers.PickupHandler{DB: db}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("🚛 Trash Pickup API Running"))
	})

	
	r.Get("/form", func(w http.ResponseWriter, r *http.Request) {
	    w.Header().Set("Content-Type", "text/html")
	    w.Write([]byte(`
	        <form method="POST" action="/pickups">
	            <input type="text" name="name" placeholder="Name" required>
	            <input type="text" name="address" placeholder="Address" required>
	            <input type="datetime-local" name="pickup_time" required>
	            <button type="submit">Schedule Pickup</button>
	        </form>
	    `))
	})
	r.Post("/pickups", handler.CreatePickup)
	r.Get("/pickups", handler.GetAllPickups)

	return r
}
