package infrastructure

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisLock struct {
	client *redis.Client
	ttl    time.Duration
}

func (r *RedisLock) AcquireLock(lockKey string) (bool, error) {
	success, err := r.client.SetNX(context.Background(), lockKey, "locked", r.ttl).Result()
	if err != nil {
		return false, err
	}
	return success, nil
}

func (r *RedisLock) ReleaseLock(lockKey string) error {
	err := r.client.Del(context.Background(), lockKey).Err()
	return err
}
