package users // data access object model

import (
	"fmt"
	"strings"

	"github.com/arya0618/my-bookstore-api/datasources/mysql/users_db"
	"github.com/arya0618/my-bookstore-api/utils/date"
	"github.com/arya0618/my-bookstore-api/utils/errors"
)

const (
	uniqueIndexEmail = "email"
	queryInsertUser  = "INSERT INTO users(first_name,last_name,email,date_created) VALUES (?,?,?,?);"
	queryGetUser     = "SELECT * from users WHERE id=? ;"
	errNoUserRow     = "no rows in result set"
)

//Get is method
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errNoUserRow) {
			return errors.NewBadRequestError(fmt.Sprintf("User with id : %d not found", user.ID))
		}
		fmt.Println(err)
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d :%s", user.ID, err.Error()))

	}
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
