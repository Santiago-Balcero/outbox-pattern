package main

import (
	"log"
	"outbox-job/db"
	"outbox-job/redis"
	"outbox-job/service"
)

func main() {
	db.Connect()
	log.Println("Pizza Orders job is running!")

	redisClient := redis.NewRedisClient("localhost:6379")

	service := service.NewService(redisClient, db.DB)

	service.ProcessPizzaCreatedOrders()
}
