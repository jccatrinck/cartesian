package env

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testEnv = `TEST=a1b2c3`

var file *os.File

func TestLoad(t *testing.T) {
	var err error

	file, err = ioutil.TempFile("", "env")
	assert.NoError(t, err)

	n, err := file.WriteString(testEnv)
	assert.NoError(t, err)
	assert.NotZero(t, n)

	err = Load(file.Name())
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	value := Get("TEST", "DEFAULT")

	assert.NotEmpty(t, value)
	assert.NotEqual(t, "DEFAULT", value)

	err := os.Remove(file.Name())
	assert.NoError(t, err)
}
