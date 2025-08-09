package service

import (
	"log"
	"outbox-job/db"
	"time"
)

func ProcessPizzaOrders() {
	for {
		var events []PizzaOrderOutbox
		err := db.DB.Where("status = ?", Pending).Find(&events).Limit(5).Error
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
			pizzaOrder, err := getPizzaOrder(event.PizzaOrderID)
			if err != nil {
				log.Println("Error fetching pizza order:", err)
				continue
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
