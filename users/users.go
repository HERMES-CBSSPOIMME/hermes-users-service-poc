package users

import (

	// 3rd Party Libs
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Uid         string `json:"userID" bson:"_id"`
	Username    string `json:"username" bson:"username"`
	Name        string `json:"name" bson:"name"`
	Surname     string `json:"surname" bson:"surname"`
	Email       string `json:"email" bson:"email"`
	Picture_URL string `json:"pictureURL" bson:"pictureURL"`
	Password    string `json:"password" bson:"password"`
}

func NewUser(username string, name string, surname string, email string, picture_url string, password string) *User {

	uuid := uuid.NewV4()

	return &User{
		Uid:         uuid.String(),
		Username:    username,
		Name:        name,
		Surname:     surname,
		Email:       email,
		Picture_URL: picture_url,
		Password:    password,
	}
}

func AssignId(user *User) error {

	uuid := uuid.NewV4()

	user.Uid = uuid.String()

	return nil
}
