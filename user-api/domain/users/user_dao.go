package users

import (
	"fmt"
	"user-api/datasources/postgresql/users_db"
	"user-api/utils/errors"
)

const (
	insertUserQuery = "INSERT INTO users_db.users(first_name, last_name, email, date_created) VALUES ($1, $2, $3, $4) RETURNING id"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	err := users_db.Client.QueryRow(insertUserQuery, user.FirstName, user.LastName, user.Email, user.DateCreated).Scan(&user.Id)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("1. error when trying to save user: %s", err.Error()))
	}

	// current := usersDB[user.Id]
	// if current != nil {
	// 	if current.Email == user.Email {
	// 		return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
	// 	}
	// 	return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	// }
	// user.DateCreated = date_utils.GetNowString()
	// usersDB[user.Id] = user
	return nil
}
