package api

import (
	"net/http"

	"github.com/codelesshub/api-generic-crud-insert/internal/custom_error"
	"github.com/codelesshub/api-generic-crud-insert/internal/service"
	"github.com/gin-gonic/gin"
)

func handleInsert(c *gin.Context) {
	collectionName := c.Param("collection")
	var data map[string]interface{}

	// Tentar fazer o bind do JSON para o map
	if err := c.BindJSON(&data); err != nil {
		HandleError(c, custom_error.NewErrInvalidPayload()) // Usando erro personalizado
		return
	}

	// Tentar inserir o payload
	id, err := service.InsertPayload(collectionName, data)
	if err != nil {
		HandleError(c, err)
		return
	}

	// Retornar o ID do documento criado com status 201 Created
	c.JSON(http.StatusCreated, gin.H{"_id": id})
}
