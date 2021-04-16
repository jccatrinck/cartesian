package services

import (
	"github.com/jccatrinck/cartesian/services/points"
	"github.com/jccatrinck/cartesian/storage"
)

// Configure all services
func Configure() (err error) {
	s := storage.Get()

	err = points.Configure(s)

	if err != nil {
		return
	}

	return
}
