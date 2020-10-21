package services

import (
	"fmt"

	"github.com/arya0618/my-bookstore-api/domain/users"
	"github.com/arya0618/my-bookstore-api/utils/errors"
)

//CreateUser is
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	// save to databse
	if err := user.Save(); err != nil {
		return nil, err
	}
	//do not have err the return pointer of user
	return &user, nil
}

//GetUser is function
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	fmt.Println("ser--", result)
	return result, nil
}
