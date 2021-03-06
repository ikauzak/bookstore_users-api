package users

import (
	"github.com/ikauzak/bookstore_users-api/utils/errors"
	"strings"
)

// dto = data transfer object

//User can
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name`
	Email       string `json:"email"`
	DateCreated string `json:"data_created"`
}

// Validate method user.Validate()
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
