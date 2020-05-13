package deep_copy_concept

import (
	"../deep_copy_concept"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReferenceCopyIssue(t *testing.T) {
	andrea := deep_copy_concept.Person{
		Name: "Andrea Mazzanti",
		Street: &deep_copy_concept.Address{
			StreetName:   "Via Panoramica",
			StreetNumber: 57,
		},
	}
	luca := andrea
	assert.EqualValues(t, andrea.Street.StreetNumber, luca.Street.StreetNumber)
	luca.Street.StreetNumber = 34
	assert.NotEqual(t, andrea.Street.StreetNumber, luca.Street.StreetNumber) // this fail by intention - reference to copied object
}

func TestPerson_DeepCopyCopy(t *testing.T) {
	andrea := deep_copy_concept.Person{
		Name: "Andrea Mazzanti",
		Street: &deep_copy_concept.Address{
			StreetName:   "Via Panoramica",
			StreetNumber: 57,
		},
	}
	luca := andrea.DeepCopy()
	assert.EqualValues(t, andrea.Street.StreetNumber, luca.Street.StreetNumber)
	luca.Street.StreetNumber = 34
	assert.NotEqual(t, andrea.Street.StreetNumber, luca.Street.StreetNumber) // this fail by intention - reference to copied object
}
