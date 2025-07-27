// Package dts handles application configuration, including loading environment variables.
package dts

type User struct {
	ID        int    `json:"id" example:"1"`
	FirstName string `json:"firstname" example:"This is the user firstname" binding:"required"`
	LastName  string `json:"lastname" example:"This is the user lastname" binding:"required"`
	Age       int    `json:"age" example:"24" binding:"required"`
}
