package main

import "fmt"

// implementing interface

var _ User = &superUser{}

type superUser struct {
	Name      string
	IsBlocked bool
}

func (s *superUser) ChangeBlockStatus(isBlock bool) {
	s.IsBlocked = isBlock
}

func (s *superUser) ChangeName(newName string) {
	s.Name = newName
}

func (s *superUser) GetName() string {
	return s.Name
}

var _ User = &user{}

type user struct {
	Name      string
	IsBlocked bool
}

func (u *user) ChangeBlockStatus(isBlock bool) {
	u.IsBlocked = isBlock
}

func (u *user) ChangeName(newName string) {
	u.Name = newName
}

func (u *user) GetName() string {
	return u.Name
}

type User interface {
	ChangeName(newName string)
	ChangeBlockStatus(isBlock bool)
	GetName() string
}

// realise User interface
func newUser(name string, isBlocked bool) User {
	u := user{Name: name, IsBlocked: isBlocked}
	//u2 := superUser{Name: name, IsBlocked: isBlocked}
	return &u
}

func newSuperUser(name string, isBlocked bool) User {
	u2 := superUser{Name: name, IsBlocked: isBlocked}
	return &u2
}

//func main() {
func interfaceExample() {
	user := newUser("John", false)
	superUser := newSuperUser("Jack", false)
	fmt.Println(user)           // &{John false}
	fmt.Println(user.GetName()) // John

	fmt.Println(superUser)           // &{Jack false}
	fmt.Println(superUser.GetName()) // Jack
}
