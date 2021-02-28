package main

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"workshop/internal/handler"
)

func main() {
	h := handler.NewHandler()
	r := chi.NewRouter()
	r.Get("/hello", h.Hello)
	log.Print("starting server")
	err := http.ListenAndServe(":8000", r)
	log.Fatal(err)
	log.Print("shutting server down")
}
