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

	r.Post("/pickups", handler.CreatePickup)
	r.Get("/pickups", handler.GetAllPickups)

	return r
}
