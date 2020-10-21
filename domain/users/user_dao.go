package users // data access object model

import (
	"fmt"

	"github.com/arya0618/my-bookstore-api/utils/errors"
)

// mocke db
var (
	userDB = make(map[int64]*User)
)

//Get is method
func (user *User) Get() *errors.RestErr {
	result := userDB[user.ID]
	if result == nil {
		//sprinf for formatting
		return errors.NewNotfoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	fmt.Println("get--", user)
	return nil
}

//Save is method to save data into db
func (user *User) Save() *errors.RestErr {
	current := userDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registred", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}
	userDB[user.ID] = user
	return nil
}