package main

import (
	"log"
	"net/http"

	"crud/config"
	"crud/handlers"
	"crud/routes"
)

func main() {
	// Carregar configurações
	cfg := config.LoadConfig()

	// Inicializar dados de exemplo
	handlers.InitializeMovies()

	// Configurar rotas
	r := routes.SetupRoutes()

	// Iniciar servidor
	log.Printf("🚀 Iniciando servidor na porta %s", cfg.Port)
	log.Printf("📖 Documentação disponível em: http://%s", cfg.GetServerAddress())
	log.Printf("🏥 Health check: http://%s/health", cfg.GetServerAddress())
	log.Printf("🎬 API de filmes: http://%s/api/v1/movies", cfg.GetServerAddress())

	log.Fatal(http.ListenAndServe(cfg.GetServerAddress(), r))
}
