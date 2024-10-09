package dragonite

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

func (o *Validate) RateLimit(ctx context.Context, config RateLimitConfig) *Validate {
	key := fmt.Sprintf("rate-limit:%s:%s", config.As, config.SubAs)
	var val string
	var valInt int64

	r := o.rds.Get(ctx, key)
	val, o.Error = r.Result()
	if o.Error != nil {
		if o.Error.Error() == redis.Nil.Error() {
			o.Error = o.rds.Set(ctx, key, 1, config.In*time.Millisecond).Err()
			if o.Error != nil {
				return o
			}
			return o
		}
		return o
	}

	valInt, o.Error = strconv.ParseInt(val, 10, 32)
	if o.Error != nil {
		return o
	}

	if valInt >= config.Limit {
		o.Error = fmt.Errorf("rate limit exceeded")
		return o
	}

	o.Error = o.rds.Incr(ctx, key).Err()
	if o.Error != nil {
		return o
	}

	return o
}
