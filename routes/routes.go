// Package routes handles application configuration, including loading environment variables.
package routes

import (
	"app/arrayvsslice"
	"app/users"
	"app/websocket"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine) {
	// Serve Swagger UI at /swagger/*any
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Add WebSocket route (e.g., /ws/:topic)
	router.GET("/ws/:topic", websocket.WebSocketHandler)

	v1 := router.Group("/api/v1")
	{
		v1.GET("", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello from Gin!")
		})

		RegisterArrayRoutes(v1)
		RegisterUserRoutes(v1)
		NotificationRoutes(v1)
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

// NotificationRoutes notification group
func NotificationRoutes(router *gin.RouterGroup) {
	handler := websocket.Notification{}

	group := router.Group("/notify")
	group.POST("/send", handler.SendNotification)
}
