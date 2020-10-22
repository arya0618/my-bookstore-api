package users // data access object model

import (
	"fmt"

	"github.com/arya0618/my-bookstore-api/datasources/mysql/users_db"
	"github.com/arya0618/my-bookstore-api/utils/date"
	"github.com/arya0618/my-bookstore-api/utils/errors"
)

// mocke db
var (
	userDB = make(map[int64]*User)
)

//Get is method
func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		fmt.Println("---in get()---", err)
		panic(err)
	}
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

	//user.DateCreated = now.Format("2006-01-02T15:04:05Z") //YYYY-MM-DD
	user.DateCreated = date.GetNowString()
	userDB[user.ID] = user
	return nil
}
