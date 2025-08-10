package main

import (
	"outbox-consumer/redis"
	"outbox-consumer/service"
)

func main() {
	redis.Connect()
	service.FetchOrderEvents()
}
