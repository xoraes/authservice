package models

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}
type UserList struct {
	Users []User `json:"users,omitempty"`
}

func (user *User) HashPassword() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	return string(bytes), err
}

func (user *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		logrus.Debugf(err.Error())
		return false
	}
	return true
}

func (user *User) Valid() bool {
	if user.Password == "" || user.Email == "" || user.FirstName == "" || user.LastName == "" {
		return false
	}
	return true
}
