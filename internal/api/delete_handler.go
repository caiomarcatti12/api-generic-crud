package api

import (
	"net/http"

	"github.com/codelesshub/api-generic-crud-insert/internal/service"
	"github.com/gin-gonic/gin"
)

func handleDelete(c *gin.Context) {
	collectionName := c.Param("collection")
	id := c.Param("id")

	// Tenta deletar o documento
	err := service.DeletePayload(collectionName, id)
	if err != nil {
		HandleError(c, err)
		return
	}

	// Se a deleção for bem-sucedida, retorna um status 204 No Content
	c.Status(http.StatusNoContent)
}
