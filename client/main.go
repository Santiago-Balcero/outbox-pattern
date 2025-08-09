package main

import (
	"log"
	"outbox-client/service"
)

const numOfOrders = 100

func main() {
	log.Println("Starting pizza order creation...")
	for i := 0; i < numOfOrders; i++ {
		service.CreatePizzaOrder()
	}
}
