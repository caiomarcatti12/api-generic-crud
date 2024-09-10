package repository

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

// SelectPayloadByID retorna o documento da coleção com o ID especificado
func FindSchemaByCollectionAndCommand(collectionName string, action string) (map[string]interface{}, error) {
	collection := mgm.CollectionByName("schemas")
	filter := bson.M{"collection": collectionName, "action": action}

	var result map[string]interface{}
	err := collection.FindOne(mgm.Ctx(), filter).Decode(&result)
	
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}
