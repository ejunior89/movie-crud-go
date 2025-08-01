package routes

import (
	"crud/handlers"
	"crud/middleware"

	"github.com/gorilla/mux"
)

// SetupRoutes configura todas as rotas da aplicação
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Aplicar middleware de logging
	r.Use(middleware.Logger)

	// Rotas da API
	api := r.PathPrefix("/api/v1").Subrouter()

	// Rotas de filmes
	api.HandleFunc("/movies", handlers.GetMovies).Methods("GET")
	api.HandleFunc("/movies/{id}", handlers.GetMovie).Methods("GET")
	api.HandleFunc("/movies", handlers.CreateMovie).Methods("POST")
	api.HandleFunc("/movies/{id}", handlers.UpdateMovie).Methods("PUT")
	api.HandleFunc("/movies/{id}", handlers.DeleteMovie).Methods("DELETE")

	// Rota de health check
	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

	// Rota raiz
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	return r
}
