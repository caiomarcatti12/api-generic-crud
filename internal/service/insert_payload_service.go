package service

import (
	"github.com/codelesshub/api-generic-crud-insert/internal/custom_error"
	"github.com/codelesshub/api-generic-crud-insert/internal/repository"
)

// InsertPayload processa e insere o payload na coleção especificada
func InsertPayload(collectionName string, data map[string]interface{}) (string, *custom_error.CustomError) {

	// Validar o payload com o schema
	errSchema := ValidatePayloadBySchema(collectionName, "INSERT", data)

	if errSchema != nil {
		return "", errSchema
	}

	id, err := repository.InsertPayload(collectionName, data)

	if err != nil {
		return "", custom_error.NewErrDatabase(err)
	}

	return id, nil
}
