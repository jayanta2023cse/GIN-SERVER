// Package websocket handles application configuration, including loading environment variables.
package websocket

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Notification struct{}

func (n Notification) SendNotification(c *gin.Context) {
	var req struct {
		Topic   string      `json:"topic"`
		Message interface{} `json:"message"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := EmitMessageToTopic(req.Topic, req.Message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "notification sent"})
}
