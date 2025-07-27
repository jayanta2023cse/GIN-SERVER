// Package routes handles application configuration, including loading environment variables.
package routes

import (
	"app/arrayvsslice"
	"app/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello from Gin!")
		})

		RegisterArrayRoutes(v1)
		RegisterUserRoutes(v1)
	}
}

// RegisterArrayRoutes array vs slice group
func RegisterArrayRoutes(router *gin.RouterGroup) {
	handler := arrayvsslice.ArrayVsSlice{}

	group := router.Group("/arrslc")
	group.GET("/modifyslice", handler.ModifyASlice)
	group.GET("/appendslice", handler.AppendToASlice)
}

// RegisterUserRoutes array vs slice group
func RegisterUserRoutes(router *gin.RouterGroup) {
	handler := users.Users{}

	group := router.Group("/user")
	group.GET("/get", handler.GetUsers)
	group.POST("/create", handler.CreateUser)
}
