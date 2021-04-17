package redis

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	err := Configure()
	assert.NoError(t, err)
	assert.NotNil(t, Client)

	ctx := context.Background()

	status := Client.Set(ctx, "TEST", "TEST", 150*time.Millisecond)
	assert.NoError(t, status.Err())

	stringCMD := Client.Get(ctx, "TEST")

	value, err := stringCMD.Result()
	assert.NoError(t, err)

	assert.Equal(t, "TEST", value)
}
