package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigure(t *testing.T) {
	err := Configure()
	assert.NoError(t, err)
	assert.NotNil(t, Get())
}
