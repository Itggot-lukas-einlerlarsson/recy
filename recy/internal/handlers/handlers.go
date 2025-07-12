package handlers

import (
    "net/http"
    "recy/internal/services"
)

var (
    recycleService *services.RecycleService
)

func InitHandlers() {
    recycleService = services.NewRecycleService()
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    // Handle home route
    w.Write([]byte("Welcome to the Recycle App!"))
}
