package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jccatrinck/cartesian/libs/env"
)

// Client is a Redis client connection
var Client *redis.Client

// Export redis lib used values
var (
	// Nil reply returned by Redis when key does not exist.
	Nil = redis.Nil
)

// Configure setups the Redis client
func Configure() (err error) {
	host := env.Get("REDIS_HOST", "localhost")
	port := env.Get("REDIS_PORT", "6379")

	Client = redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
	})

	ctx := context.Background()

	_, err = Client.Ping(ctx).Result()

	if err != nil {
		return
	}

	return
}
