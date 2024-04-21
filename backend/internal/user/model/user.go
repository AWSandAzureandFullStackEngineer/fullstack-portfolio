package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
}

func NewUser(firstName, lastName, username, email, password string) (*User, error) {
	newUser := &User{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Email:     email,
		Password:  password,
	}
	return newUser, nil
}
