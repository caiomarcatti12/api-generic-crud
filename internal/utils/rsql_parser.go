package utils

import (
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

// ConvertRSQLToMongoFilter converte uma string RSQL para um filtro MongoDB
func ConvertRSQLToMongoFilter(rsql string) (bson.M, error) {
	filter := bson.M{}

	// Exemplo básico: Assume que a string RSQL é "campo==valor;outroCampo==outroValor"
	conditions := strings.Split(rsql, ";")
	for _, condition := range conditions {
		parts := strings.Split(condition, "==")
		if len(parts) == 2 {
			field := parts[0]
			value := parts[1]

			// Verifica se o valor é numérico e remove a comparação regex, caso seja
			if valueInt, err := strconv.Atoi(value); err == nil {
				// Valor é um número, faça uma comparação direta
				filter[field] = valueInt
			} else {
				// Adiciona um filtro regex para ignorar o tipo do campo e realizar a comparação case-insensitive
				filter[field] = bson.M{
					"$regex":   "^" + value + "$", // Ajuste para correspondência exata
					"$options": "i",
				}
			}
		}
	}

	// Aqui você pode adicionar suporte para operadores como >, <, >=, etc.

	return filter, nil
}
