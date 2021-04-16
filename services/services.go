package services

import (
	"github.com/jccatrinck/cartesian/services/points"
)

// Configure all services
func Configure() (err error) {
	err = points.Configure()

	if err != nil {
		return
	}

	return
}
