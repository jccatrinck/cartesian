package points

import (
	"context"
	"os"

	"github.com/jccatrinck/cartesian/libs/env"
	"github.com/jccatrinck/cartesian/services/points/model"
)

var storage Storage

func loadStorage(s Storage) {
	storage = s
}

func loadPoints() (err error) {
	filepath := env.Get("POINTS_FILE", "data/points.json")

	file, err := os.Open(filepath)

	if err != nil {
		return
	}

	err = storage.LoadPoints(file)

	if err != nil {
		return
	}

	return
}

// Configure Points service
func Configure(s Storage) (err error) {
	loadStorage(s)

	err = loadPoints()

	if err != nil {
		return
	}

	return
}

// GetPointsByDistance using storage
func GetPointsByDistance(ctx context.Context, point model.Point, distance int) (points []model.RelativePoint, err error) {
	points, err = storage.GetPointsByDistance(point, distance)

	if err != nil {
		return
	}

	return
}
