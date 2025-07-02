package config

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func loadRedis(redisHost, redisPort string) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisHost + ":" + redisPort,
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return redisClient, nil
}
