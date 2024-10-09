package dragonite

import (
	"github.com/go-redis/redis/v8"
)

func NewValidate(rds *redis.Client) *Validate {
	return &Validate{
		rds: rds,
	}
}

type Validate struct {
	rds   *redis.Client
	Error error
}
