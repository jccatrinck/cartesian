package model

import (
	"math/big"
	"sort"
)

// Point coordinates
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Distance of another point using Manhattan
func (p1 Point) Distance(p2 Point) int {
	x1 := big.NewInt(int64(p1.X))
	x2 := big.NewInt(int64(p2.X))

	y1 := big.NewInt(int64(p1.Y))
	y2 := big.NewInt(int64(p2.Y))

	diffX := big.NewInt(0).Sub(x1, x2)
	diffX.Abs(diffX)

	diffY := big.NewInt(0).Sub(y1, y2)
	diffY.Abs(diffY)

	distance := big.NewInt(0).Add(diffX, diffY)

	return int(distance.Int64())
}

// RelativePoint to other point
type RelativePoint struct {
	Point
	Distance int `json:"distance"`
}

// RelativePointSortedList keeps a list of sorted points
type RelativePointSortedList struct {
	points []RelativePoint
}

// Add sorted data into list
func (r *RelativePointSortedList) Add(p RelativePoint) {
	n := len(r.points)

	index := sort.Search(n, func(i int) bool {
		return r.points[i].Distance > p.Distance
	})

	r.points = append(r.points, RelativePoint{})
	copy(r.points[index+1:], r.points[index:])
	r.points[index] = p
}

// Get sorted list
func (r *RelativePointSortedList) Get() []RelativePoint {
	return r.points
}
