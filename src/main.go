package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/josewenzel/flagify/src/api/route"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Could not load .env file: %s", err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.HandleFunc("/status", route.StatusRequest())

	http.ListenAndServe(":8123", router)
}
