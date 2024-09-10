package service

import (
	"github.com/codelesshub/api-generic-crud-insert/internal/custom_error"
	"github.com/codelesshub/api-generic-crud-insert/internal/repository"
	"github.com/codelesshub/api-generic-crud-insert/internal/utils"
)

// SelectPayloadWithFilter processa a consulta filtrada e retorna os dados e o total de registros
func SelectPayloadWithFilter(collectionName string, rsql string, limit int64, offset int64) (map[string]interface{}, *custom_error.CustomError) {
	// Converter RSQL para filtro MongoDB
	filter, err := utils.ConvertRSQLToMongoFilter(rsql)
	if err != nil {
		return nil, custom_error.NewErrDatabase(err)
	}

	// Chamar o repositório passando o filtro já parseado
	data, totalCount, err := repository.SelectPayloadWithFilter(collectionName, filter, limit, offset)
	if err != nil {
		return nil, custom_error.NewErrDatabase(err)
	}

	// Montar o resultado final
	result := map[string]interface{}{
		"data":       data,
		"totalCount": totalCount,
	}

	return result, nil
}
