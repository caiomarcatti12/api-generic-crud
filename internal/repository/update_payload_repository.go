package repository

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

// UpdatePayload atualiza o documento na coleção com o ID especificado
func UpdatePayload(collectionName string, id string, updateData map[string]interface{}) error {
	collection := mgm.CollectionByName(collectionName)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updateData}

	_, err := collection.UpdateOne(mgm.Ctx(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
