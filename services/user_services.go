package services

import (
	"github.com/arya0618/my-bookstore-api/domain/users"
)

//CreateUser is
func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
