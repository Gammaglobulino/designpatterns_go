package factory_generator

import (
	"../factory_generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactoryEmployee(t *testing.T) {
	developerFactory := factory_generator.NewFactoryEmployee("developer", 60000)
	managerFactory := factory_generator.NewFactoryEmployee("manager", 80000)

	developer := developerFactory("Andrea Mazzanti")
	manager := managerFactory("Luca Fuso")

	assert.EqualValues(t, 60000, developer.AnnualIncome)
	assert.EqualValues(t, 80000, manager.AnnualIncome)
}
