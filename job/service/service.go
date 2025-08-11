package service

import (
	"encoding/json"
	"fmt"
	"log"
	"outbox-job/db"
	"time"

	"gorm.io/gorm"
)

type Service struct {
	Messenger Messenger
	DB        *gorm.DB
}

func NewService(messenger Messenger, db *gorm.DB) *Service {
	return &Service{
		Messenger: messenger,
		DB:        db,
	}
}

func (s *Service) ProcessPizzaCreatedOrders() {
	for {
		var events []PizzaOrderOutbox
		err := s.DB.
			Where("status = ? AND event_type = ?", Pending, PizzaOrderCreated).
			Order("created_at ASC").
			Limit(10).
			Find(&events).
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

			if err := s.Messenger.Ping(); err != nil {
				log.Println(fmt.Sprintf("Error pinging %s:", s.Messenger.GetName()), err)
				break
			}

			if err := s.Messenger.SendMessage("pizza-orders", event.Payload); err != nil {
				log.Println(fmt.Sprintf("Error publishing pizza order to %s:", s.Messenger.GetName()), err)
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
