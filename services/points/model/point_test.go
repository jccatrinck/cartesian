package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	point := Point{
		X: 9,
		Y: -9,
	}

	data, err := json.Marshal(point)

	assert.NoError(t, err)
	assert.NotEmpty(t, data)

	expected := `{"x":9,"y":-9}`
	assert.JSONEq(t, expected, string(data))

	d := point.Distance(Point{X: 1, Y: 1})
	assert.Equal(t, 18, d)
}

func TestRelativePointSortedList(t *testing.T) {
	sortedList := RelativePointSortedList{}

	sortedList.Add(RelativePoint{Distance: 3})
	sortedList.Add(RelativePoint{Distance: 2})
	sortedList.Add(RelativePoint{Distance: 1})

	list := sortedList.Get()

	assert.Equal(t, 1, list[0].Distance)
	assert.Equal(t, 2, list[1].Distance)
	assert.Equal(t, 3, list[2].Distance)
}
