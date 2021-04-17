package points

import (
	"io"

	"github.com/jccatrinck/cartesian/services/points/model"
)

// Storage needs definition
type Storage interface {
	LoadPoints(reader io.ReadSeeker) error
	GetPointsByDistance(point model.Point, distance int) ([]model.RelativePoint, error)
}
