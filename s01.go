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

func (user User) SetFirstName(string) {

}

func (user User) SetLastName(string) {

}

func (user User) FullName() string {
	return ""
}

func ResetUser(input User) {

}

func IsUser(input User) bool {
	return true
}

func ProcessUser(input UserInterface) string {
	return ""
}
