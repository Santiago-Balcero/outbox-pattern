package main

import (
	"log"
	"outbox-job/db"
	"outbox-job/redis"
	"outbox-job/service"
)

func main() {
	db.Connect()
	redis.Connect()
	log.Println("Pizza Orders job is running!")

	service.ProcessPizzaCreatedOrders()
}
