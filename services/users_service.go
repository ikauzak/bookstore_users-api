package services

import (
	"github.com/ikauzak/bookstore_users-api/domain/users"
	"github.com/ikauzak/bookstore_users-api/utils/errors"
)

//CreateUser does
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return &user, nil
}
