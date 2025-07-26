// Package helpers handles application configuration, including loading environment variables.
package helpers

import (
	"github.com/gin-gonic/gin"
)

var blankObject = struct{}{}

func RenderJSON(c *gin.Context, status int, data interface{}, message string, errMsg string, ok bool) {
	c.JSON(status, gin.H{
		"error":   errMsg,
		"status":  ok,
		"data":    data,
		"message": message,
	})
}

func HandleError(c *gin.Context, status int, err interface{}, message string) {
	response := map[string]interface{}{
		"status":  false,
		"data":    blankObject,
		"message": message,
	}

	switch e := err.(type) {
	case map[string]string:
		response["error"] = e
	case error:
		response["error"] = e.Error()
	default:
		response["error"] = "Unknown error occurred"
	}

	c.JSON(status, response)
}
