package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ikauzak/bookstore_users-api/domain/users"
	"io/ioutil"
	"net/http"
)

//GetUser does
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

//CreateUser does
func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println(err)
	fmt.Println(string(bytes))
	c.String(http.StatusNotImplemented, "implement me")
}

//SearchUser does
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
