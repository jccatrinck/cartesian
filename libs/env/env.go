package env

import (
	"os"

	"github.com/joho/godotenv"
)

var env map[string]string

// Load env from file
func Load(filename string) (err error) {
	return godotenv.Load(filename)
}

// Get a env value by its key otherwise return default value
func Get(key string, def string) (v string) {
	v, exists := os.LookupEnv(key)

	if !exists {
		return def
	}

	return
}
