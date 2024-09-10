package api

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter configura as rotas do servidor
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/:collection", handleInsert)
	router.DELETE("/:collection/:id", handleDelete)
	router.PUT("/:collection/:id", handleUpdate)
	router.GET("/:collection/:id", handleSelectByID)
	router.GET("/:collection", handleSelectWithFilter)
	router.PATCH("/:collection/:id/:action", handleAction)

	return router
}
