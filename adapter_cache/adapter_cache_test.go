package adapter_cache

import (
	"../adapter_cache"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testImage struct {
	image []adapter_cache.Point
}

var img = testImage{[]adapter_cache.Point{
	{X: 0, Y: 0}, {X: 5, Y: 5},
},
}

func (i *testImage) GetPoints() []adapter_cache.Point {
	return i.image
}

func TestNewRectangle(t *testing.T) {
	rc := adapter_cache.NewRectangle(6, 4)
	assert.EqualValues(t, 0, rc.Lines[0].X1)
	assert.EqualValues(t, 0, rc.Lines[0].Y1)
	assert.EqualValues(t, 5, rc.Lines[0].X2)
	assert.EqualValues(t, 0, rc.Lines[0].Y2)

	assert.EqualValues(t, 0, rc.Lines[3].X1)
	assert.EqualValues(t, 3, rc.Lines[3].Y1)
	assert.EqualValues(t, 5, rc.Lines[3].X2)
	assert.EqualValues(t, 3, rc.Lines[3].Y2)

}

func TestDrawPoints(t *testing.T) {
	assert.EqualValues(t, "*     \n      \n      \n      \n      \n     *\n", adapter_cache.DrawPoints(&img))
	fmt.Println(adapter_cache.DrawPoints(&img))
}
func TestFromLineToPoint(t *testing.T) {
	line := adapter_cache.Line{
		X1: 0,
		Y1: 0,
		X2: 0,
		Y2: 3,
	}
	adapter := adapter_cache.VectorToRasterAdapter{}
	adapter.AddLine(line)
	assert.EqualValues(t, 4, len(adapter.Points))
	assert.EqualValues(t, "*\n*\n*\n*\n", adapter_cache.DrawPoints(&adapter))

}
