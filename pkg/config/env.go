package config

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv carrega as variáveis de ambiente do arquivo .env
func LoadEnv() {
	err := godotenv.Load("../../configs/.env")
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
}
