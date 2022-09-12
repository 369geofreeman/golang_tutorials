package main

import "log"

type Database interface {
	GetUser() string
	GetAllUsers() []string
}

type DefaultDatabase struct{}

func (db DefaultDatabase) GetUser() string {
	return "bob"
}

type FamousDatabase struct{}

func (db FamousDatabase) GetUser() string {
	return "Will Smith"
}

func (db FamousDatabase) GetAllUsers() []string {
	return []string{}
}

type Greeter interface {
	Greet(userName string)
}

type NiceGreeter struct{}

func (ng NiceGreeter) Greet(userName string) {
	log.Printf("Hello %s!!! Nice to meet you", userName)
}

type Program struct {
	Db      Database
	Greeter Greeter
}

func (p Program) Excecute() {
	user := p.Db.GetUser()
	p.Greeter.Greet(user)
}

func main() {
	// Injection
	// db := DefaultDatabase{}
	db := FamousDatabase{}
	greeter := NiceGreeter{}

	p := Program{
		Db:      db,
		Greeter: greeter,
	}

	p.Excecute()
}
