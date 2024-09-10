package repository

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

// DeletePayload remove o documento da coleção com o ID especificado
func DeletePayload(collectionName string, id string) error {
	collection := mgm.CollectionByName(collectionName)
	filter := bson.M{"_id": id}

	_, err := collection.DeleteOne(mgm.Ctx(), filter)
	if err != nil {
		return err
	}

	return nil
}
