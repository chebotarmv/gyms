package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	"github.com/chebotarmv/gymshark/server/handlers"
)

func main() {
	r := chi.NewRouter()
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(corsMiddleware.Handler)

	r.Get("/calculate", handlers.CalculatePacks)
	http.ListenAndServe(":8033", r)
}
