package functional_factory_generator

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

func NewEmployeeFactory(position string, annualincome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualincome}
	}
}
