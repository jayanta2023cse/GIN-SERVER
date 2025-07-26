// Package routes handles application configuration, including loading environment variables.
package routes

import (
	"app/arrayvsslice"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from Gin!")
	})

	RegisterArrayRoutes(router)
}

// RegisterArrayRoutes array vs slice group
func RegisterArrayRoutes(router *gin.Engine) {
	handler := arrayvsslice.ArrayVsSlice{}

	group := router.Group("/arrslc")
	group.GET("/modifyslice", handler.ModifyASlice)
	group.GET("/appendslice", handler.AppendToASlice)
}
