package service

import (
	"context"
	"log"
	"outbox-consumer/redis"

	redisclient "github.com/redis/go-redis/v9"
)

func FetchOrderEvents() {
	if err := redis.Client.Ping(context.Background()).Err(); err != nil {
		log.Println("Error pinging Redis:", err)
		return
	}

	subscription := redis.Client.Subscribe(context.Background(), "pizza-orders")
	messagesch := subscription.Channel()

	var msg *redisclient.Message
	for {
		msg = <-messagesch
		println("Received message:", msg.Payload, "from topic:", msg.Channel)
	}
}
