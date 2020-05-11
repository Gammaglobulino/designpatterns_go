package builder_facets

import (
	"../builder_facets"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPersonBuilder(t *testing.T) {
	pb := builder_facets.NewPersonBuilder()
	pb.
		Lives().
		At("123 London Road").
		In("London").
		WithPostcode("SW12BC").
		Works().
		At("Fabrikam").
		AsA("Programmer").
		Earning(123000)
	person := pb.Build()
	assert.EqualValues(t, "123 London Road", person.StreetAddress)
}
