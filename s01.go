package structs

import "reflect"

type User struct {
	firstName string
	lastName  string
}

type UserInterface interface {
	SetFirstName(string)
	SetLastName(string)
	FullName() string
}

func (user *User) SetFirstName(firstName string) {
	user.firstName = firstName
}

func (user *User) SetLastName(lastName string) {
	user.lastName = lastName
}

func (user *User) FullName() string {
	return user.lastName + " " + user.firstName
}

func New() User {
	return User{
		firstName: "",
		lastName:  "",
	}
}

func ResetUser(input interface{}) {
	if user, ok := input.(*User); ok {
		user.firstName = ""
		user.lastName = ""
	}
}

func IsUser(input interface{}) bool {
	user := User{}
	if reflect.TypeOf(user) != reflect.TypeOf(input) {
		return false
	}
	return true
}

func ProcessUser(input UserInterface) string {
	input.SetFirstName("")
	input.SetLastName("")
	return input.FullName()
}
