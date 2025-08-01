package config

import (
	"log"
	"os"
	"strconv"
)

// Config armazena as configurações da aplicação
type Config struct {
	Port    string
	Host    string
	LogFile string
	Debug   bool
}

// LoadConfig carrega as configurações do ambiente
func LoadConfig() *Config {
	config := &Config{
		Port:    getEnv("PORT", "8000"),
		Host:    getEnv("HOST", "localhost"),
		LogFile: getEnv("LOG_FILE", ""),
		Debug:   getEnvAsBool("DEBUG", false),
	}

	log.Printf("Configurações carregadas: Port=%s, Host=%s, Debug=%t",
		config.Port, config.Host, config.Debug)

	return config
}

// getEnv obtém uma variável de ambiente ou retorna um valor padrão
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsBool obtém uma variável de ambiente como boolean
func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}

// GetServerAddress retorna o endereço completo do servidor
func (c *Config) GetServerAddress() string {
	return c.Host + ":" + c.Port
}
