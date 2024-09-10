package api

import (
	"net/http"

	"github.com/codelesshub/api-generic-crud-insert/internal/custom_error"
	"github.com/gin-gonic/gin"
)

// HandleError mapeia os erros para os códigos HTTP corretos e responde ao cliente
func HandleError(c *gin.Context, customErr *custom_error.CustomError) {
	// Verifica o tipo de erro
	switch customErr.Type {
	case "NotFound":
		c.JSON(http.StatusNotFound, gin.H{"error": customErr.Message})
	case "InvalidID", "InvalidPayload":
		c.JSON(http.StatusBadRequest, gin.H{"error": customErr.Message})
	case "DatabaseError":
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao acessar o banco de dados"})
	case "InvalidSchema", "InvalidSchemaCompile":
		c.JSON(http.StatusBadRequest, gin.H{"error": customErr.Message})
	case "InvalidFilter":
		c.JSON(http.StatusBadRequest, gin.H{"error": customErr.Message})

	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro desconhecido"})
	}
}
