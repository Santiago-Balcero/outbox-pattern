package redis

import "github.com/redis/go-redis/v9"

var Client *redis.Client

func Connect() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})
}
