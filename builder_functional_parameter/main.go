package builder_functional_parameter

type Person struct {
	Name, Position string
}

type personMod func(*Person)
type PersonBuilder struct {
	actions []personMod
}

func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(person *Person) {
		person.Name = name
	})
	return b

}

func (b *PersonBuilder) WorkAs(position string) *PersonBuilder {
	b.actions = append(b.actions, func(person *Person) {
		person.Position = position
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
