package service

import (
	"github.com/codelesshub/api-generic-crud-insert/internal/custom_error"
	"github.com/codelesshub/api-generic-crud-insert/internal/repository"
)

// SelectPayloadByID retorna o documento da coleção com o ID especificado
func SelectPayloadByID(collectionName string, id string) (map[string]interface{}, *custom_error.CustomError) {
	// Se necessário, adicionar aqui alguma lógica de negócios, validação, transformação, etc.
	result, err := repository.SelectPayloadByID(collectionName, id)
	if err != nil {
		return nil, custom_error.NewErrDatabase(err)
	}

	// Verificar se o documento foi encontrado
	if result == nil {
		return nil, custom_error.NewErrNotFound()
	}

	return result, nil
}
