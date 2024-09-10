package repository

import (
	"github.com/google/uuid"
	"github.com/kamva/mgm/v3"
)

// InsertPayload insere o payload na coleção e retorna o UUID gerado
func InsertPayload(collectionName string, data map[string]interface{}) (string, error) {
	id := uuid.New().String()

	data["_id"] = id

	collection := mgm.CollectionByName(collectionName)
	_, err := collection.InsertOne(mgm.Ctx(), data)
	if err != nil {
		return "", err
	}

	return id, nil
}
