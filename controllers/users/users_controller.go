package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ikauzak/bookstore_users-api/domain/users"
	"github.com/ikauzak/bookstore_users-api/services"
	"net/http"
)

//GetUser does
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

//CreateUser does
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		//TODO: return bad request to the caller.
		return
	}
	result, saveErr := services.CreateUser(user)

	if saveErr != nil {
		//TODO: Handle user creation error
		return
	}

	c.JSON(http.StatusCreated, result)
}

//SearchUser does
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
