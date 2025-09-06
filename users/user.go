// Package users handles application configuration, including loading environment variables.
package users

import (
	"app/dts"
	"net/http"

	"github.com/gin-gonic/gin"
)

var listOfUsers []dts.User
var lastID int

type Users struct{}

// GetUsers godoc
// @Summary List all users
// @Description Retrieve a list of all user items
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} dts.User "List of users"
// @Router /api/v1/user/get [get]
func (e Users) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, listOfUsers)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Add a new user item to the list
// @Tags users
// @Accept json
// @Produce json
// @Param user body dts.User true "User item to create"
// @Success 201 {object} dts.User "Created user item"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Router /api/v1/user/create [post]
func (e Users) CreateUser(c *gin.Context) {
	var user dts.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lastID++
	user.ID = lastID
	listOfUsers = append(listOfUsers, user)

	// messageKey := constants.UserModule
	// messageValue := fmt.Sprintf("New user created: %+v", user) // message content
	// if err := kafka.PublishMessage(messageKey, messageValue); err != nil {
	// 	log.Printf("Failed to send Kafka message: %v", err)
	// 	// optional: you can still return success to client
	// }

	c.JSON(http.StatusCreated, user)
}
