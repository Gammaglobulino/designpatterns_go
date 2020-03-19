package builder_functional_parameter

import (
	"../builder_functional_parameter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPersonBuilder_Called(t *testing.T) {
	b := builder_functional_parameter.PersonBuilder{}
	p := b.Called("Andrea").Build()
	assert.EqualValues(t, "Andrea", p.Name)
}

func TestPersonBuilder_WorkAs(t *testing.T) {
	b := builder_functional_parameter.PersonBuilder{}
	p := b.WorkAs("Programmer").Build()
	assert.EqualValues(t, "Programmer", p.Position)
}

func TestBuilder(t *testing.T) {
	b := builder_functional_parameter.PersonBuilder{}
	p := b.Called("Andrea").WorkAs("Programmer").Build()
	assert.EqualValues(t, "Programmer", p.Position)
	assert.EqualValues(t, "Andrea", p.Name)
}
