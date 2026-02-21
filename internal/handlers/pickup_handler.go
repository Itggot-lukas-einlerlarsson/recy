package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"trash-pickup-app/internal/models"
)

type PickupHandler struct {
	DB *sql.DB
}

func (h *PickupHandler) CreatePickup(w http.ResponseWriter, r *http.Request) {
	var pickup models.Pickup

	if err := json.NewDecoder(r.Body).Decode(&pickup); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO pickups (name, address, pickup_time) VALUES ($1, $2, $3)`
	_, err := h.DB.Exec(query, pickup.Name, pickup.Address, pickup.PickupTime)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Pickup scheduled successfully!",
	})
}


func (h *PickupHandler) GetAllPickups(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("SELECT id, name, address, pickup_time FROM pickups ORDER BY id DESC")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var pickups []models.Pickup

	for rows.Next() {
		var pickup models.Pickup
		err := rows.Scan(&pickup.ID, &pickup.Name, &pickup.Address, &pickup.PickupTime)
		if err != nil {
			http.Error(w, "Error reading data", http.StatusInternalServerError)
			return
		}
		pickups = append(pickups, pickup)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pickups)
}
