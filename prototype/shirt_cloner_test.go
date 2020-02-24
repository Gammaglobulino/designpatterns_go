package prototype

import (
	"../prototype"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetShirtCloner(t *testing.T) {
	shirtcache := prototype.GetShirtsCloner()
	if shirtcache == nil {
		t.Fatal("received cache was nil")
	}
	item1, err := shirtcache.GetClone(prototype.White)
	if err != nil {
		t.Error(err)
	}
	assert.EqualValues(t, item1, prototype.WhiteProto)
}
func TestShirtClone(t *testing.T) {
	shirtcache := prototype.GetShirtsCloner()
	item1, _ := shirtcache.GetClone(prototype.White)
	_, ok := item1.(*prototype.Shirt)
	assert.True(t, ok, "Type assertion failed")
}
func TestShirt_GetInfo(t *testing.T) {
	shirtcache := prototype.GetShirtsCloner()
	item1, _ := shirtcache.GetClone(prototype.White)
	assert.EqualValues(t, item1.GetInfo(), "Price 15.00 , SKU empty , color 0")
}
