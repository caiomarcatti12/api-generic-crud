package service

import (
	"encoding/json"

	"github.com/codelesshub/api-generic-crud-insert/internal/custom_error"
	"github.com/codelesshub/api-generic-crud-insert/internal/repository"
	"github.com/xeipuuv/gojsonschema"
)

// InsertPayload processa e insere o payload na coleção especificada
func ValidatePayloadBySchema(collectionName string, command string, data map[string]interface{}) *custom_error.CustomError {
	schemaDB, err := repository.FindSchemaByCollectionAndCommand(collectionName, command)
	if err != nil {
		return custom_error.NewErrDatabase(err)
	}

	if schemaDB == nil {
		return nil
	}

	// Convertendo o map para JSON string
	schemaData, err := json.Marshal(schemaDB["schema"])
	if err != nil {
		return custom_error.NewErrInvalidSchemaCompile(err)
	}

	schemaLoader := gojsonschema.NewStringLoader(string(schemaData))
	documentLoader := gojsonschema.NewGoLoader(data)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return custom_error.NewErrInvalidSchemaCompile(err)
	}

	if !result.Valid() {
		return custom_error.NewErrInvalidSchema(result.Errors())
	}

	return nil
}
