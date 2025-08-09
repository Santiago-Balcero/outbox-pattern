package service

import "gorm.io/gorm"

type Status string

const (
	Pending   Status = "pending"
	Completed Status = "completed"
	Failed    Status = "failed"
)

type PizzaOrder struct {
	gorm.Model
	Flavor   string  `json:"flavor"`
	Size     string  `json:"size"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	Address  string  `json:"address"`
	UserName string  `json:"user_name"`
}

type PizzaOrderOutbox struct {
	gorm.Model
	PizzaOrderID uint   `json:"pizza_order_id"`
	Status       Status `json:"status"`
	Error        string `json:"error"`
}
