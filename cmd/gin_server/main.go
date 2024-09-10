package main

import (
	"log"

	"github.com/codelesshub/api-generic-crud-insert/internal/api"
	"github.com/codelesshub/api-generic-crud-insert/pkg/config"
	"github.com/codelesshub/api-generic-crud-insert/pkg/database"
)

func main() {
	// Carrega as variáveis de ambiente do arquivo .env
	config.LoadEnv()

	// Inicializa a conexão com o banco de dados
	database.InitDB()

	// Configura o router
	router := api.SetupRouter()

	// Inicia o servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Falha ao iniciar o servidor: ", err)
	}
}
