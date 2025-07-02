package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func loadRedis(redisHost, redisPort string) (*redis.Client, error) {
	redis := redis.NewClient(&redis.Options{
		Addr: redisHost + ":" + redisPort,
	})

	if err := redis.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("error in pinging redis: %s", err.Error())
	}

	return redis, nil
}
