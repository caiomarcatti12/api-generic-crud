package api

import (
	"net/http"

	"github.com/codelesshub/api-generic-crud-insert/internal/custom_error"
	"github.com/codelesshub/api-generic-crud-insert/internal/service"
	"github.com/gin-gonic/gin"
)

func handleSelectByID(c *gin.Context) {
	collectionName := c.Param("collection")
	id := c.Param("id")

	// Chama o serviço para buscar o documento pelo ID
	result, err := service.SelectPayloadByID(collectionName, id)
	if err != nil {
		HandleError(c, err)
		return
	}

	// Se o documento não for encontrado, retorna um erro 404
	if result == nil {
		HandleError(c, custom_error.NewErrNotFound())
		return
	}

	// Retorna o documento encontrado com status 200
	c.JSON(http.StatusOK, result)
}
