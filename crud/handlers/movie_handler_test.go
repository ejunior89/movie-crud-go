package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"crud/models"
	"github.com/gorilla/mux"
)

func TestGetMovies(t *testing.T) {
	// Limpar dados de teste
	movies = []models.Movie{}
	InitializeMovies()

	// Criar requisição
	req, err := http.NewRequest("GET", "/api/v1/movies", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Criar response recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetMovies)

	// Executar requisição
	handler.ServeHTTP(rr, req)

	// Verificar status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verificar se retornou JSON válido
	var response models.MovieResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("response is not valid JSON: %v", err)
	}

	// Verificar se success é true
	if !response.Success {
		t.Errorf("expected success to be true, got %v", response.Success)
	}
}

func TestCreateMovie(t *testing.T) {
	// Limpar dados de teste
	movies = []models.Movie{}

	// Dados de teste
	movieData := models.MovieRequest{
		Isbn:        "123456",
		Title:       "Test Movie",
		Director:    &models.Director{Firstname: "John", Lastname: "Doe"},
		ReleaseYear: 2023,
		Genre:       "Ação",
		Rating:      8.5,
	}

	// Converter para JSON
	jsonData, _ := json.Marshal(movieData)

	// Criar requisição
	req, err := http.NewRequest("POST", "/api/v1/movies", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Criar response recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateMovie)

	// Executar requisição
	handler.ServeHTTP(rr, req)

	// Verificar status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verificar se retornou JSON válido
	var response models.MovieResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("response is not valid JSON: %v", err)
	}

	// Verificar se success é true
	if !response.Success {
		t.Errorf("expected success to be true, got %v", response.Success)
	}

	// Verificar se o filme foi criado
	if len(movies) != 1 {
		t.Errorf("expected 1 movie, got %d", len(movies))
	}
}

func TestGetMovie(t *testing.T) {
	// Limpar dados de teste
	movies = []models.Movie{}
	InitializeMovies()

	// Criar requisição
	req, err := http.NewRequest("GET", "/api/v1/movies/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Criar router para capturar parâmetros
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/movies/{id}", GetMovie).Methods("GET")

	// Criar response recorder
	rr := httptest.NewRecorder()

	// Executar requisição
	router.ServeHTTP(rr, req)

	// Verificar status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verificar se retornou JSON válido
	var response models.MovieResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("response is not valid JSON: %v", err)
	}

	// Verificar se success é true
	if !response.Success {
		t.Errorf("expected success to be true, got %v", response.Success)
	}
}

func TestHealthCheck(t *testing.T) {
	// Criar requisição
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Criar response recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheck)

	// Executar requisição
	handler.ServeHTTP(rr, req)

	// Verificar status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verificar se retornou JSON válido
	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("response is not valid JSON: %v", err)
	}

	// Verificar se status é "OK"
	if response["status"] != "OK" {
		t.Errorf("expected status to be 'OK', got %v", response["status"])
	}
} 