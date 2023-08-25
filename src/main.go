package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/josewenzel/flagify/src/api/route"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.HandleFunc("/status", route.StatusRequest())

	http.ListenAndServe(":8123", router)
}
