package users

import (
	"fmt"
	"github.com/ikauzak/bookstore_users-api/utils/errors"
	"time"
)

var (
	usersDB = make(map[int64]*User)
)

// dao = data access object

//Get user
func (user *User) Get() *errors.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

//Save saves the
func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}
	now := time.Now().UTC()
	user.DateCreated = now.Format("02-01-2006T15:04:05Z")

	usersDB[user.ID] = user
	return nil
}
