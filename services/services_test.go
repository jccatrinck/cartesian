package services

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/jccatrinck/cartesian/storage"
	"github.com/stretchr/testify/assert"
)

func TestConfigure(t *testing.T) {
	file, err := ioutil.TempFile("", "dummy")
	assert.NoError(t, err)

	n, err := file.WriteString("[]")
	assert.NoError(t, err)
	assert.NotZero(t, n)

	err = os.Setenv("POINTS_FILE", file.Name())
	assert.NoError(t, err)

	err = storage.Configure()
	assert.NoError(t, err)

	err = Configure()
	assert.NoError(t, err)
}
