package factory_generator

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

func NewFactoryEmployee(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{
			Name:         name,
			Position:     position,
			AnnualIncome: annualIncome,
		}
	}
}
