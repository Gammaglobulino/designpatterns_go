package builder

import (
	"../builder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestManufacturingDirector_SetBuilderCar(t *testing.T) {
	manufactoringComplex := builder.ManufacturingDirector{}
	carbuilder := &builder.CarBuilder{}
	manufactoringComplex.SetBuilder(carbuilder)
	manufactoringComplex.Construct()
	car := carbuilder.Build()
	assert.EqualValues(t, 4, car.Wheels)
}

func TestManufacturingDirector_SetBuilderBike(t *testing.T) {
	manufactoringComplex := builder.ManufacturingDirector{}
	bikebuilder := &builder.BikeBuilder{}
	manufactoringComplex.SetBuilder(bikebuilder)
	manufactoringComplex.Construct()
	bike := bikebuilder.Build()
	assert.EqualValues(t, 2, bike.Wheels)
}
