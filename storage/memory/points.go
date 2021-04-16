package memory

import (
	"io"
	"sort"

	"github.com/jccatrinck/cartesian/services/points"
	"github.com/jccatrinck/cartesian/services/points/model"
)

type memoryPoints struct {
	points []model.Point
	xAxis  []int
	yAxis  []int
}

// LoadPoints implements points.Storage interface
func (m *memoryPoints) LoadPoints(reader io.ReadSeeker) (err error) {
	pw, err := points.NewWalker(reader)

	if err != nil {
		return
	}

	err = pw.Run(func(chunk []model.Point) (err error) {
		for _, point := range chunk {
			m.points = append(m.points, point)
		}

		return
	})

	if err != nil {
		return
	}

	m.xAxis = make([]int, 0, len(m.points))
	for i := range m.points {
		m.xAxis = append(m.xAxis, i)
	}
	sort.Slice(m.xAxis, func(i, j int) bool {
		a := m.xAxis[i]
		b := m.xAxis[j]

		return m.points[a].X < m.points[b].X
	})

	m.yAxis = make([]int, 0, len(m.points))
	for i := range m.points {
		m.yAxis = append(m.yAxis, i)
	}
	sort.Slice(m.yAxis, func(i, j int) bool {
		a := m.yAxis[i]
		b := m.yAxis[j]

		return m.points[a].Y < m.points[b].Y
	})

	return
}

// GetPointsByDistance implements points.Storage interface
func (m *memoryPoints) GetPointsByDistance(pointA model.Point, distance int) (relativePoints []model.RelativePoint, err error) {
	sortedList := model.RelativePointSortedList{}

	// Get nearest point relative to X as start point
	xAxisIndex, ok := divideAndConquer(m.xAxis, pointA.X, 0, func(i int) int {
		return m.points[i].X
	})

	// Taxicab circle
	xAxisDiameter := map[int]struct{}{}

	if ok {
		// Taxicab circle radius - down
		for i := xAxisIndex - 1; i > 0; i-- {
			pointIndex := m.xAxis[i]

			distanceX := pointA.X - m.points[pointIndex].X

			if distanceX > distance {
				break
			}

			xAxisDiameter[pointIndex] = struct{}{}
		}

		// Taxicab circle radius - up
		for i := xAxisIndex; i < len(m.xAxis); i++ {
			pointIndex := m.xAxis[i]

			distanceX := m.points[pointIndex].X - pointA.X

			if distanceX > distance {
				break
			}

			xAxisDiameter[pointIndex] = struct{}{}
		}
	}

	// Get nearest point relative to Y as start point
	yAxisIndex, ok := divideAndConquer(m.yAxis, pointA.Y, 0, func(i int) int {
		return m.points[i].Y
	})

	if ok {
		// Taxicab circle radius - left
		for i := yAxisIndex - 1; i > 0; i-- {
			pointIndex := m.yAxis[i]
			pointB := m.points[pointIndex]

			distanceY := pointB.Y - pointA.Y

			if distanceY > distance {
				break
			}

			if _, exists := xAxisDiameter[pointIndex]; exists {
				pointsDistance := pointA.Distance(pointB)

				if pointsDistance <= distance {
					// Add point already sorted
					sortedList.Add(model.RelativePoint{
						Point:    pointB,
						Distance: pointsDistance,
					})
				}
			}
		}

		// Taxicab circle radius - right
		for i := yAxisIndex; i < len(m.yAxis); i++ {
			pointIndex := m.yAxis[i]
			pointB := m.points[pointIndex]

			distanceY := pointA.Y - pointB.Y

			if distanceY > distance {
				break
			}

			if _, exists := xAxisDiameter[pointIndex]; exists {
				pointsDistance := pointA.Distance(pointB)

				if pointsDistance <= distance {
					// Add point already sorted
					sortedList.Add(model.RelativePoint{
						Point:    pointB,
						Distance: pointsDistance,
					})
				}
			}
		}
	}

	relativePoints = sortedList.Get()

	return
}

func divideAndConquer(items []int, item, acc int, getter func(int) int) (int, bool) {
	total := len(items)

	if total == 1 {
		return 0, true
	} else if total == 0 {
		return -1, false
	}

	middle := int(total / 2)

	if getter(items[middle]) == item {
		return middle, true
	}

	// Get diff of each side of slice
	left := item - getter(items[middle-1])
	right := getter(items[middle]) - item

	// Find by proximity, not exact value. The lower the closer
	if left < right {
		i, ok := divideAndConquer(items[:middle], item, acc, getter)
		return i + acc, ok
	}

	i, ok := divideAndConquer(items[middle:], item, acc, getter)
	return i + acc + middle, ok
}
