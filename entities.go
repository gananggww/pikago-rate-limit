package dragonite

import (
	"time"
)

type RateLimitConfig struct {
	Limit int64
	In    time.Duration
	SubAs string
	As    string
}
