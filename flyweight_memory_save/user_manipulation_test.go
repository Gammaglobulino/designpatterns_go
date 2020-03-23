package flyweight_memory_save

import (
	"../flyweight_memory_save"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	andrea := flyweight_memory_save.NewUser("Andrea")
	assert.EqualValues(t, "Andrea", andrea.FullName)
}

func TestNewUser2(t *testing.T) {
	andrea := flyweight_memory_save.NewUser2("Andrea")
	stefano := flyweight_memory_save.NewUser2("Stefano")

	assert.EqualValues(t, "\x00", andrea.String())
	assert.EqualValues(t, "\x01", stefano.String())
}
