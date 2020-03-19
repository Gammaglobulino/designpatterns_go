package functional_factory_generator

import (
	"../functional_factory_generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEmployeeFactory(t *testing.T) {
	developerFactory := functional_factory_generator.NewEmployeeFactory("Developer", 60000)
	managerFactory := functional_factory_generator.NewEmployeeFactory("Manager", 80000)

	developer := developerFactory("Andrea")
	manager := managerFactory("Stefano")

	assert.EqualValues(t, "Andrea", developer.Name)
	assert.EqualValues(t, "Stefano", manager.Name)
}
