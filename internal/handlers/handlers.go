
package handlers

import (
    "net/http"
    services "recy-app/internal/service"
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
