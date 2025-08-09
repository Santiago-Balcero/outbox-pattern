package service

import (
	"encoding/json"
	"log"
	"outbox-job/db"
	"time"
)

func ProcessPizzaCreatedOrders() {
	for {
		var events []PizzaOrderOutbox
		err := db.DB.Where("status = ? AND event_type = ?", Pending, PizzaOrderCreated).Find(&events).Limit(5).Error
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

			if event.Payload == "" {
				pizzaOrder, err = getPizzaOrder(event.PizzaOrderID)
				if err != nil {
					log.Println("Error fetching pizza order:", err)
					continue
				}
			} else {
				err = json.Unmarshal([]byte(event.Payload), &pizzaOrder)
				if err != nil {
					log.Println("Error unmarshalling pizza order:", err)
					continue
				}
			}

			getPayment(pizzaOrder)
			deliver(pizzaOrder)

			err = db.DB.Model(&PizzaOrderOutbox{}).Where("pizza_order_id = ?", pizzaOrder.ID).Update("status", Completed).Error
			if err != nil {
				log.Println("Error updating outbox status for pizza order:", pizzaOrder.ID, "-", err)
				continue
			}

			log.Println("Pizza order processed! 🍕")
		}
	}
}

func getPizzaOrder(id uint) (PizzaOrder, error) {
	var pizzaOrder PizzaOrder
	err := db.DB.Where("id = ?", id).First(&pizzaOrder).Error
	return pizzaOrder, err
}

func getPayment(pizzaOrder PizzaOrder) {
	log.Println("Payment processing $:", pizzaOrder.Price)
}

func deliver(pizzaOrder PizzaOrder) {
	log.Println("Delivering pizza to:", pizzaOrder.UserName, "-", pizzaOrder.Address)
}
