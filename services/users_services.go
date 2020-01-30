package services

import (
	"github.com/brendankoral/go-books-users-api/domain/users"
	"github.com/brendankoral/go-books-users-api/utils/errors"
	"net/http"
	"os"
)

func CreateUser(user User) (*users.User, *errors.RestErr) {
	return &user, nil
}
