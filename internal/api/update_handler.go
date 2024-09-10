package api

import (
	"net/http"

	"github.com/codelesshub/api-generic-crud-insert/internal/custom_error"
	"github.com/codelesshub/api-generic-crud-insert/internal/service"
	"github.com/gin-gonic/gin"
)

func handleUpdate(c *gin.Context) {
	collectionName := c.Param("collection")
	id := c.Param("id")

	var data map[string]interface{}

	// Tentar fazer o bind do JSON para o map
	if err := c.BindJSON(&data); err != nil {
		HandleError(c, custom_error.NewErrInvalidPayload()) // Usando erro personalizado
		return
	}

	err := service.UpdatePayload(collectionName, id, data)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
