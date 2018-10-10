package user

import (

	// 3rd Party Libs
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Uid         string
	Username    string
	Name        string
	Surname     string
	Email       string
	Picture_URL string
}

func NewUser(username string, name string, surname string, email string, picture_url string) *User {

	uuid := uuid.NewV4()

	return &User{
		Uid:         uuid.String(),
		Username:    username,
		Name:        name,
		Surname:     surname,
		Email:       email,
		Picture_URL: picture_url,
	}
}
