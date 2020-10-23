package users // data access object model

import (
	"fmt"
	"strings"

	"github.com/arya0618/my-bookstore-api/datasources/mysql/users_db"
	"github.com/arya0618/my-bookstore-api/utils/date"
	"github.com/arya0618/my-bookstore-api/utils/errors"
)

// mocke db
var (
	userDB = make(map[int64]*User)
)

const (
	uniqueIndexEmail = "email"
	queryInsertUser  = "INSERT INTO users(first_name,last_name,email,date_created) VALUES (?,?,?,?);"
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
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date.GetNowString()
	insertResult, insertErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if insertErr != nil {
		if strings.Contains(insertErr.Error(), uniqueIndexEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save the user :%s", insertErr.Error()))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save the user :%s", insertErr.Error()))
	}
	user.ID = userId

	// current := userDB[user.ID]
	// if current != nil {
	// 	if current.Email == user.Email {
	// 		return errors.NewBadRequestError(fmt.Sprintf("email %s already registred", user.Email))
	// 	}
	// 	return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	// }

	// //user.DateCreated = now.Format("2006-01-02T15:04:05Z") //YYYY-MM-DD
	// user.DateCreated = date.GetNowString()
	// userDB[user.ID] = user
	return nil
}
