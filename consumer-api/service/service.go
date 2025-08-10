package service

import (
	"context"
	"outbox-consumer/redis"

	redisclient "github.com/redis/go-redis/v9"
)

func FetchOrderEvents() {
	subscription := redis.Client.Subscribe(context.Background(), "pizza-orders")
	messagesch := subscription.Channel()

	var msg *redisclient.Message
	for {
		msg = <-messagesch
		println("Received message:", msg.Payload, "from topic:", msg.Channel)
	}
}
