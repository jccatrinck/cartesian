package points

import (
	"context"

	"github.com/jccatrinck/cartesian/services/points/model"
)

// Configure Points service
func Configure() (err error) {
	return
}

// GetPointsByDistance using storage
func GetPointsByDistance(ctx context.Context, point model.Point, distance int) (points []model.RelativePoint, err error) {
	return
}
