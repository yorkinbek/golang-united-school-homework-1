package structs

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
	return user.firstName + " " + user.lastName
}

func NewUser() *User {
	return &User{}
}

func ResetUser(input interface{}) {
	if user, ok := input.(*User); ok {
		user.firstName = ""
		user.lastName = ""
	}
}

func IsUser(input interface{}) bool {
	_, ok := input.(*User)
	return ok
}

func ProcessUser(input UserInterface) string {
	input.SetFirstName(NewUser().firstName)
	input.SetLastName(NewUser().lastName)
	return input.FullName()
}
