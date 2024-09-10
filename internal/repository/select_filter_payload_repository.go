package repository

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SelectPayloadWithFilter realiza a consulta filtrada e retorna os dados e o total de registros
func SelectPayloadWithFilter(collectionName string, filter bson.M, limit int64, offset int64) ([]bson.M, int64, error) {
	collection := mgm.CollectionByName(collectionName)

	// Obter o total de registros
	totalCount, err := collection.CountDocuments(mgm.Ctx(), filter)
	if err != nil {
		return []bson.M{}, 0, err
	}

	// Realizar a consulta com limite e offset
	findOptions := options.Find().SetLimit(limit).SetSkip(offset)
	cursor, err := collection.Find(mgm.Ctx(), filter, findOptions)
	if err != nil {
		return []bson.M{}, 0, err
	}
	defer cursor.Close(mgm.Ctx())

	var data []bson.M
	if err = cursor.All(mgm.Ctx(), &data); err != nil {
		return []bson.M{}, 0, err
	}

	if data == nil {
		data = []bson.M{}
	}
	return data, totalCount, nil
}
