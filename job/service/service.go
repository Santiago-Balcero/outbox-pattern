package service

import (
	"context"
	"encoding/json"
	"log"
	"outbox-job/db"
	"outbox-job/redis"
	"time"
)

func ProcessPizzaCreatedOrders() {
	for {
		var events []PizzaOrderOutbox
		err := db.DB.
			Where("status = ? AND event_type = ?", Pending, PizzaOrderCreated).
			Order("created_at ASC").
			Find(&events).
			Limit(10).
			Error
		if err != nil {
			log.Println("Error fetching pizza orders:", err)
			continue
		}

		if len(events) == 0 {
			log.Println("No pending pizza orders to process.")
			time.Sleep(1 * time.Second) // Wait before checking again
			continue
		}

		for _, event := range events {
			log.Println("Processing pizza order:", event.PizzaOrderID)

			var pizzaOrder PizzaOrder
			err = json.Unmarshal([]byte(event.Payload), &pizzaOrder)
			if err != nil {
				log.Println("Error unmarshalling pizza order:", err)
				continue
			}

			if err := redis.Client.Ping(context.Background()).Err(); err != nil {
				log.Println("Error pinging Redis:", err)
				break
			}

			if err := redis.Client.Publish(context.Background(), "pizza-orders", event.Payload).Err(); err != nil {
				log.Println("Error publishing pizza order to Redis:", err)
				break
			}

			err = db.DB.Model(&PizzaOrderOutbox{}).Where("pizza_order_id = ?", pizzaOrder.ID).Update("status", Completed).Error
			if err != nil {
				log.Println("Error updating outbox status for pizza order:", pizzaOrder.ID, "-", err)
				continue
			}

			log.Println("Pizza order processed! 🍕")
		}
	}
}
