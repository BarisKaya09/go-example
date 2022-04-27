package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	name     string
	username string
	age      int
	lang     []string
}

func (e *Employee) updateName(name string) {
	e.name = name
}

func (e *Employee) updateUsername(username string) {
	e.username = username
}

func (e *Employee) updateAge(age int) {
	e.age = age
}

func (e *Employee) updateLang(lang string) {
	e.lang = append(e.lang, lang)
}

func main() {
	Employee := Employee{"baris", "bariskaya", 18, []string{"Python", "JavaScript"}}
	Employee.updateName("burak")
	Employee.updateUsername("burakkaya")
	Employee.updateAge(26)
	Employee.updateLang("Go")

	fmt.Println(Employee.name, Employee.username, Employee.age, strings.Join(Employee.lang, " "))
}