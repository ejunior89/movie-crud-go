package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"crud/models"

	"github.com/gorilla/mux"
)

var movies []models.Movie

// GetMovies retorna todos os filmes
func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := models.MovieResponse{
		Success: true,
		Message: "Filmes recuperados com sucesso",
		Data:    movies,
	}

	json.NewEncoder(w).Encode(response)
}

// GetMovie retorna um filme específico por ID
func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			response := models.MovieResponse{
				Success: true,
				Message: "Filme encontrado",
				Data:    item,
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	// Filme não encontrado
	w.WriteHeader(http.StatusNotFound)
	response := models.MovieResponse{
		Success: false,
		Error:   "Filme não encontrado",
	}
	json.NewEncoder(w).Encode(response)
}

// CreateMovie cria um novo filme
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movieRequest models.MovieRequest
	err := json.NewDecoder(r.Body).Decode(&movieRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := models.MovieResponse{
			Success: false,
			Error:   "Formato JSON inválido",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Validar campos obrigatórios
	if movieRequest.Isbn == "" || movieRequest.Title == "" || movieRequest.Director == nil {
		w.WriteHeader(http.StatusBadRequest)
		response := models.MovieResponse{
			Success: false,
			Error:   "Campos obrigatórios: isbn, title e director são necessários",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Criar novo filme
	movie := models.Movie{
		ID:          strconv.FormatInt(time.Now().UnixNano(), 10),
		Isbn:        movieRequest.Isbn,
		Title:       movieRequest.Title,
		Director:    movieRequest.Director,
		ReleaseYear: movieRequest.ReleaseYear,
		Genre:       movieRequest.Genre,
		Rating:      movieRequest.Rating,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, movie)

	response := models.MovieResponse{
		Success: true,
		Message: "Filme criado com sucesso",
		Data:    movie,
	}
	json.NewEncoder(w).Encode(response)
}

// UpdateMovie atualiza um filme existente
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range movies {
		if item.ID == params["id"] {
			var movieRequest models.MovieRequest
			err := json.NewDecoder(r.Body).Decode(&movieRequest)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				response := models.MovieResponse{
					Success: false,
					Error:   "Formato JSON inválido",
				}
				json.NewEncoder(w).Encode(response)
				return
			}

			// Validar campos obrigatórios
			if movieRequest.Isbn == "" || movieRequest.Title == "" || movieRequest.Director == nil {
				w.WriteHeader(http.StatusBadRequest)
				response := models.MovieResponse{
					Success: false,
					Error:   "Campos obrigatórios: isbn, title e director são necessários",
				}
				json.NewEncoder(w).Encode(response)
				return
			}

			// Atualizar filme
			movies[i].Isbn = movieRequest.Isbn
			movies[i].Title = movieRequest.Title
			movies[i].Director = movieRequest.Director
			movies[i].ReleaseYear = movieRequest.ReleaseYear
			movies[i].Genre = movieRequest.Genre
			movies[i].Rating = movieRequest.Rating
			movies[i].UpdatedAt = time.Now()

			response := models.MovieResponse{
				Success: true,
				Message: "Filme atualizado com sucesso",
				Data:    movies[i],
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	// Filme não encontrado
	w.WriteHeader(http.StatusNotFound)
	response := models.MovieResponse{
		Success: false,
		Error:   "Filme não encontrado",
	}
	json.NewEncoder(w).Encode(response)
}

// DeleteMovie remove um filme
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)

			response := models.MovieResponse{
				Success: true,
				Message: "Filme removido com sucesso",
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	// Filme não encontrado
	w.WriteHeader(http.StatusNotFound)
	response := models.MovieResponse{
		Success: false,
		Error:   "Filme não encontrado",
	}
	json.NewEncoder(w).Encode(response)
}

// HealthCheck verifica se a API está funcionando
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"status":    "OK",
		"message":   "API está funcionando",
		"timestamp": time.Now().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(response)
}

// HomeHandler retorna informações sobre a API
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"message": "Bem-vindo à API de Filmes em Go",
		"version": "1.0.0",
		"endpoints": map[string]string{
			"health":      "/health",
			"movies":      "/api/v1/movies",
			"movie_by_id": "/api/v1/movies/{id}",
		},
		"documentation": "Consulte o README.md para mais informações",
	}
	json.NewEncoder(w).Encode(response)
}

// InitializeMovies inicializa alguns filmes de exemplo
func InitializeMovies() {
	movies = append(movies, models.Movie{
		ID:          "1",
		Isbn:        "438227",
		Title:       "Movie One",
		Director:    &models.Director{Firstname: "John", Lastname: "Doe"},
		ReleaseYear: 2020,
		Genre:       "Ação",
		Rating:      8.5,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	movies = append(movies, models.Movie{
		ID:          "2",
		Isbn:        "45455",
		Title:       "Movie Two",
		Director:    &models.Director{Firstname: "Steve", Lastname: "Smith"},
		ReleaseYear: 2021,
		Genre:       "Drama",
		Rating:      7.8,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
}
