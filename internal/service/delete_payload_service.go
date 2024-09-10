package service

import (
	"github.com/codelesshub/api-generic-crud-insert/internal/custom_error"
	"github.com/codelesshub/api-generic-crud-insert/internal/repository"
)

// DeletePayload remove o documento da coleção com o ID especificado
func DeletePayload(collectionName string, id string) *custom_error.CustomError {
	// Verificar se o documento existe
	existingDocument, err := repository.SelectPayloadByID(collectionName, id)
	if err != nil {
		return custom_error.NewErrDatabase(err)
	}

	if existingDocument == nil {
		return custom_error.NewErrNotFound()
	}

	// Se o documento existir, prosseguir com a deleção
	err = repository.DeletePayload(collectionName, id)
	if err != nil {
		return custom_error.NewErrDatabase(err)
	}
	return nil
}
