package service

import (
	"github.com/codelesshub/api-generic-crud-insert/internal/custom_error"
	"github.com/codelesshub/api-generic-crud-insert/internal/repository"
)

// UpdatePayload processa e atualiza o documento na coleção especificada
func UpdatePayload(collectionName string, id string, updateData map[string]interface{}) *custom_error.CustomError {

	errSchema := ValidatePayloadBySchema(collectionName, "UPDATE", updateData)

	if errSchema != nil {
		return errSchema
	}

	// Verificar se o documento existe
	existingDocument, err := repository.SelectPayloadByID(collectionName, id)
	if err != nil {
		return custom_error.NewErrDatabase(err)
	}

	if existingDocument == nil {
		return custom_error.NewErrNotFound()
	}

	// Se o documento existir, prosseguir com a atualização
	err = repository.UpdatePayload(collectionName, id, updateData)
	if err != nil {
		return custom_error.NewErrDatabase(err)
	}

	return nil
}
