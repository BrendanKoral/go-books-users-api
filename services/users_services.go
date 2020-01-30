package services

import "github.com/brendankoral/go-books-users-api/domain/users"

func CreateUser(user User) (*users.User, error) {
	return &user, nil
}
