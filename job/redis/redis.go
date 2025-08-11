package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

type RedisClient struct {
	Name   string
	Client *redis.Client
}

func NewRedisClient(url string) *RedisClient {
	return &RedisClient{
		Name: "redis",
		Client: redis.NewClient(&redis.Options{
			Addr: url,
		}),
	}
}

func (r *RedisClient) GetName() string {
	return r.Name
}

func (r *RedisClient) Ping() error {
	return r.Client.Ping(context.Background()).Err()
}

func (r *RedisClient) SendMessage(channel string, message string) error {
	return r.Client.Publish(context.Background(), channel, message).Err()
}
