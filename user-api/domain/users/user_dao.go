package users

import (
	"database/sql"
	"fmt"
	"strings"
	"user-api/datasources/postgresql/users_db"
	"user-api/utils/date_utils"
	"user-api/utils/errors"
)

const (
	indexUniqueEmail = "users_email_key"
	insertUserQuery  = "INSERT INTO users_db.users(first_name, last_name, email, date_created) VALUES ($1, $2, $3, $4) RETURNING id"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users_db.users WHERE id = $1"
	queryUpdateUser  = "UPDATE users_db.users SET first_name=$1, last_name=$2, email=$3 WHERE id=$4"
)

// var (
// 	usersDB = make(map[int64]*User)
// )

func (user *User) Get() *errors.RestErr {
	rows := users_db.Client.QueryRow(queryGetUser, user.Id)
	err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
		} else {
			return errors.NewInternalServerError("internal server error trying to find the user")
		}
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	user.DateCreated = date_utils.GetNowString()

	err := users_db.Client.QueryRow(insertUserQuery, user.FirstName, user.LastName, user.Email, user.DateCreated).Scan(&user.Id)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
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

func (user *User) Update() *errors.RestErr {
	_, err := users_db.Client.Exec(queryUpdateUser, user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error updating user %d - %s", user.Id, err.Error()))
	}

	return nil
}
