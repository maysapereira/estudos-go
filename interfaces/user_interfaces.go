package interfaces

import "fmt"

type User struct {
	ID   string
	Nome string
}

type UserMethods interface {
	GetUserById(id string)
	InsertUser(name string)
}

func (u *User) GetUserById(id string)  {}
func (u *User) InsertUser(name string) {}

func test() {
	u := &User{}
	var userMethods UserMethods

	userMethods = u
	fmt.Println(userMethods)
}
