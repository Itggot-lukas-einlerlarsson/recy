package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"trash-pickup-app/internal/db"
	"trash-pickup-app/internal/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database := db.NewPostgresConnection()

	router := routes.SetupRoutes(database)

	port := os.Getenv("PORT")

	log.Println("Server running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
