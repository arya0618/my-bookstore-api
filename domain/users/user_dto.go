package users

import (
	"strings"

	"github.com/arya0618/my-bookstore-api/utils/errors"
)

//User is
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

//Validate is method to validate user info
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(user.Email)
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email")
	}
	return nil
}
