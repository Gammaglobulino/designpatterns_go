package singleton

import (
	"../singleton"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInstance(t *testing.T) {
	counter := singleton.GetInstance()
	assert.NotNil(t, counter)
}
func TestSingleton_AddOne(t *testing.T) {
	counter := singleton.GetInstance()
	assert.NotNil(t, counter)
	assert.EqualValues(t, 1, counter.AddOne())
}
