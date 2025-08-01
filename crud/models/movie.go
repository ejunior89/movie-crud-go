package models

import (
	"time"
)

// Movie representa um filme no sistema
type Movie struct {
	ID          string    `json:"id"`
	Isbn        string    `json:"isbn" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Director    *Director `json:"director" validate:"required"`
	ReleaseYear int       `json:"release_year"`
	Genre       string    `json:"genre"`
	Rating      float64   `json:"rating"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Director representa o diretor de um filme
type Director struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
}

// MovieRequest representa os dados necessários para criar/atualizar um filme
type MovieRequest struct {
	Isbn        string    `json:"isbn" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Director    *Director `json:"director" validate:"required"`
	ReleaseYear int       `json:"release_year"`
	Genre       string    `json:"genre"`
	Rating      float64   `json:"rating"`
}

// MovieResponse representa a resposta da API
type MovieResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}
