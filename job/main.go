package main

import (
	"log"
	"outbox-job/db"
	"outbox-job/service"
)

func main() {
	db.Connect()
	log.Println("Pizza Orders job is running!")

	service.ProcessPizzaOrders()
}
