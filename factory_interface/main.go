package factory_interface

import (
	"fmt"
	"strconv"
)

type Person interface {
	SayHello()
	String() string
}
type person struct {
	name string
	age  int
}

type tiredPerson struct {
	person
}

func (p *person) SayHello() {
	fmt.Printf("Hello my name is %s", p.name)
}
func (p *person) String() string {
	return p.name + strconv.Itoa(p.age)
}

func (p *tiredPerson) SayHello() {
	fmt.Printf("Hello I'm tired and my name is %s", p.name)
}
func (p *tiredPerson) String() string {
	return "Tired" + p.name + strconv.Itoa(p.age)
}

func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{person{
			name: name,
			age:  age,
		}}

	}
	return &person{
		name: name,
		age:  age,
	}
}
