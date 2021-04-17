package storage

import (
	"errors"
	"strings"

	"github.com/jccatrinck/cartesian/libs/env"
	"github.com/jccatrinck/cartesian/services/points"
	"github.com/jccatrinck/cartesian/storage/memory"
	"github.com/jccatrinck/cartesian/storage/mysql"
)

// Storage interface definitions
type Storage interface {
	points.Storage
}

// Type of storage system
type Type string

const (
	// Memory storage Type
	Memory Type = "memory"
	// MySQL storage Type
	MySQL Type = "mysql"
)

var instance Storage

// Configure storage instance based on API_STORAGE_TYPE
func Configure() (err error) {
	storageType := env.Get("API_STORAGE_TYPE", "memory")

	switch Type(strings.ToLower(storageType)) {
	case Memory:
		instance = memory.New()
	case MySQL:
		instance, err = mysql.New()

		if err != nil {
			return
		}
	default:
		err = errors.New("Invalid API_STORAGE_TYPE env variable value")
		return
	}

	return
}

// Get storage instance
func Get() Storage {
	return instance
}
