package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger é um middleware que registra informações sobre as requisições HTTP
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Criar um response writer customizado para capturar o status code
		responseWriter := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Chamar o próximo handler
		next.ServeHTTP(responseWriter, r)

		// Calcular duração da requisição
		duration := time.Since(start)

		// Log das informações da requisição
		log.Printf(
			"[%s] %s %s - %d - %v",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			responseWriter.statusCode,
			duration,
		)
	})
}

// responseWriter é um wrapper para http.ResponseWriter que captura o status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader sobrescreve o método WriteHeader para capturar o status code
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Write sobrescreve o método Write para manter a interface
func (rw *responseWriter) Write(b []byte) (int, error) {
	return rw.ResponseWriter.Write(b)
}
