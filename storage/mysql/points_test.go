package mysql

import (
	"math"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/jccatrinck/cartesian/services/points/model"
	"github.com/jccatrinck/cartesian/storage/mysql/statements"
	"github.com/stretchr/testify/assert"
)

const bulkInsertSize = 1000

const fromX, toX, fromY, toY = -180, 180, -180, 180

func BenchmarkGetPointsByDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		point := generatePoint()
		distance := randomRange(1, 5)

		_, err := m.GetPointsByDistance(point, distance)

		if err != nil {
			b.Error(err)
			return
		}
	}
}

func TestGetPointsByDistance(t *testing.T) {
	point := model.Point{X: 0, Y: 0}
	distance := 1000

	points, err := m.GetPointsByDistance(point, distance)

	assert.NoError(t, err)
	assert.NotEmpty(t, points)
}

func TestGetPointsByDistanceEmpty(t *testing.T) {
	point := model.Point{X: fromX - 1, Y: fromY - 1}
	distance := 0

	points, err := m.GetPointsByDistance(point, distance)

	assert.NoError(t, err)
	assert.Empty(t, points)
}

func generatePoints(amount int) (points []model.Point) {
	for i := 0; i < amount; i++ {
		points = append(points, generatePoint())
	}

	return
}

func generatePoint() model.Point {
	return model.Point{
		X: randomRange(fromX, toX),
		Y: randomRange(fromY, toY),
	}
}

func setupPoints() (err error) {
	total := 0

	err = m.db.Select(&total, `
		SELECT COUNT(*)
		FROM point
	`)

	if err != nil {
		return
	}

	points := generatePoints(1000000 - total)

	// Bulk insert by bulkInsertSize
	for i := 1; i <= int(math.Ceil(float64(len(points))/bulkInsertSize)); i++ {
		_, err = m.db.NamedExec(statements.LoadPoints, points[i*bulkInsertSize-bulkInsertSize:i*bulkInsertSize])

		// Ignore duplicates
		if mySQLError, ok := err.(*mysql.MySQLError); ok && mySQLError.Number == errDuplicate {
			err = nil
		}

		if err != nil {
			return
		}
	}

	return
}
