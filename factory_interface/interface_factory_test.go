package factory_interface

import (
	"../factory_interface"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPerson(t *testing.T) {
	p := factory_interface.NewPerson("Andrea", 50)
	assert.EqualValues(t, "Andrea50", p.String())
}

func TestNewTiredPerson(t *testing.T) {
	p := factory_interface.NewPerson("Andrea", 150)
	assert.EqualValues(t, "TiredAndrea150", p.String())
}
