package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("I'm alive, what is my purpose?"))
	})

	http.ListenAndServe(":3000", r)
}
