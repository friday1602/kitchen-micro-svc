package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("HTTP_PORT")

	r := chi.NewRouter()

	handler := NewHandler()
	handler.registerRoutes(r)

	log.Println("listening to port", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)

	}
}
