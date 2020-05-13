package builder_functional

import "strings"

type Person struct {
	name, position string
	telephone      string
	mail           string
}

type personMod func(*Person)
type PersonBuilder struct {
	actions []personMod
}

func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}

func (b *PersonBuilder) ContactHimUsing(telephone, mail string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.telephone = telephone
		if !strings.Contains(mail, "@") {
			panic("mail is in the wrong format")
		}
		p.mail = mail
	})
	return b
}

func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

// extend PersonBuilder
func (b *PersonBuilder) WorksAsA(position string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})
	return b
}
