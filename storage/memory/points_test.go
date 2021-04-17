package memory

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/jccatrinck/cartesian/services/points/model"
	"github.com/stretchr/testify/assert"
)

const fromX, toX, fromY, toY = -180, 180, -180, 180

const data string = `[
  {
    "x": 63,
    "y": -72
  },
  {
    "x": -94,
    "y": 89
  },
  {
    "x": -30,
    "y": -38
	}
]`

func setupPoints() (err error) {
	r := strings.NewReader(data)

	err = m.LoadPoints(r)

	if err != nil {
		return
	}

	return
}

func generatePoint() model.Point {
	return model.Point{
		X: randomRange(fromX, toX),
		Y: randomRange(fromY, toY),
	}
}

func generatePoints(amount int) (points []model.Point) {
	for i := 0; i < amount; i++ {
		points = append(points, generatePoint())
	}

	return
}

func prepareBenchmark() (m Memory, err error) {
	m = Memory{}

	points := generatePoints(100000)

	points = append(points, model.Point{X: -94, Y: 89})

	jsonBytes, err := json.Marshal(&points)

	if err != nil {
		return
	}

	err = m.LoadPoints(bytes.NewReader(jsonBytes))

	if err != nil {
		return
	}

	return
}

func BenchmarkPure(b *testing.B) {
	m, err := prepareBenchmark()

	if err != nil {
		b.Error(err)
		return
	}

	amountDistance := 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pointA := model.Point{X: -94, Y: 89}
		distance := randomRange(0, 50)

		amountDistance = amountDistance + distance

		points := []model.RelativePoint{}
		for _, pointB := range m.points {
			pointsDistance := pointA.Distance(pointB)

			if pointsDistance <= distance {
				points = append(points, model.RelativePoint{
					Point:    pointB,
					Distance: pointsDistance,
				})
			}
		}
		_ = points
	}
	b.StopTimer()
	b.ReportMetric(float64(amountDistance)/float64(b.N), "distance/op")
}

func BenchmarkGetPointsByDistance(b *testing.B) {
	m, err := prepareBenchmark()

	if err != nil {
		b.Error(err)
		return
	}

	amountDistance := 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pointA := model.Point{X: -94, Y: 89}
		distance := randomRange(0, 50)

		amountDistance = amountDistance + distance

		points, err := m.GetPointsByDistance(pointA, distance)
		if err != nil {
			b.Error()
			return
		}
		_ = points
	}
	b.StopTimer()
	b.ReportMetric(float64(amountDistance)/float64(b.N), "distance/op")
}

func TestGetPointsByDistance(t *testing.T) {
	point := model.Point{X: -94, Y: 89}
	distance := 0

	points, err := m.GetPointsByDistance(point, distance)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(points))
	assert.NotZero(t, points[0])
	assert.EqualValues(t, point, points[0].Point)
}
