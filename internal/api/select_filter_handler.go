package api

import (
	"net/http"
	"strconv"

	"github.com/codelesshub/api-generic-crud-insert/internal/service"
	"github.com/gin-gonic/gin"
)

func handleSelectWithFilter(c *gin.Context) {
	collectionName := c.Param("collection")
	rsql := c.Query("filter")

	limit, err := ParseQueryParamInt64(c, "limit", "10")
	if err != nil {
		return // O erro já foi tratado dentro de ParseQueryParamInt64
	}

	offset, err := ParseQueryParamInt64(c, "offset", "0")
	if err != nil {
		return // O erro já foi tratado dentro de ParseQueryParamInt64
	}

	result, customErr := service.SelectPayloadWithFilter(collectionName, rsql, limit, offset)
	if customErr != nil {
		HandleError(c, customErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

// ParseQueryParamInt64 tenta converter um parâmetro de consulta para int64
func ParseQueryParamInt64(c *gin.Context, paramName string, defaultValue string) (int64, error) {
	valueStr := c.DefaultQuery(paramName, defaultValue)
	value, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetro '" + paramName + "' inválido"})
		return 0, err
	}
	return value, nil
}
