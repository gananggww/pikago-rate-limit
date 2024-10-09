package dragonite

import (
	"context"
	"os"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestValidate_RateLimit(t *testing.T) {
	redisAddr := os.Getenv("REDIS_ADDR") // Get Redis address from environment variable
	if redisAddr == "" {
		redisAddr = "localhost:6379" // Default to localhost if not set
	}
	rds := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	if _, err := rds.Ping(context.Background()).Result(); err != nil {
		t.Fatalf("Error connecting to Redis: %v", err)
	}

	val := &Validate{
		rds: rds,
	}
	t.Run("Rate limit not exceeded", func(t *testing.T) {

		config := RateLimitConfig{
			SubAs: "SUBAS",
			In:    100,
			As:    "AS",
			Limit: 10,
		}

		err := val.RateLimit(context.Background(), config).Error
		assert.NoError(t, err)
	})
}
