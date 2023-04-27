package redis_lock

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

// A RedisLock is a redis lock.
type RedisLock struct {
	client     *redis.Client
	expiration time.Duration
	key        string
}

// NewRedisLock returns a RedisLock.
func NewRedisLock(client *redis.Client, key string, expiration time.Duration) *RedisLock {
	return &RedisLock{
		client:     client,
		key:        key,
		expiration: expiration,
	}
}

// Acquire acquires the lock.
func (rl *RedisLock) Acquire() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ok, err := rl.client.SetNX(ctx, rl.key, 1, rl.expiration).Result()
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("操作频繁")
	}

	return nil
}

// Release releases the lock.
func (rl *RedisLock) Release() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return rl.client.Del(ctx, rl.key).Err()
}

func (rl *RedisLock) AcquireWithTimeOut(ctx context.Context, timeOut time.Duration) (bool, error) {
	if ok, _ := rl.client.SetNX(ctx, rl.key, 1, rl.expiration).Result(); ok {
		return true, nil
	}

	time.Sleep(timeOut)
	return rl.client.SetNX(ctx, rl.key, 1, rl.expiration).Result()
}
