package database

import (
	"log"
	"os"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitDB inicializa a conexão com o banco de dados
func InitDB() {
	err := mgm.SetDefaultConfig(nil, os.Getenv("MONGO_DB"), options.Client().ApplyURI(getMongoURI()))
	if err != nil {
		log.Fatal("Falha ao conectar ao MongoDB: ", err)
	}
}

// getMongoURI monta a URI de conexão com o MongoDB a partir das variáveis de ambiente
func getMongoURI() string {
	user := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")

	if user == "" || password == "" || host == "" {
		log.Fatal("Por favor, configure as variáveis de ambiente MONGO_USER, MONGO_PASSWORD, e MONGO_HOST")
	}

	return "mongodb://" + user + ":" + password + "@" + host
}
