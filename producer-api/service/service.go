package service

import (
	"encoding/json"
	"log"
	"net/http"
	"outbox/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePizza(c *gin.Context) {
	var pizzaOrder PizzaOrder
	if err := c.ShouldBindJSON(&pizzaOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := db.DB.Begin()
	err := tx.Create(&pizzaOrder).Error
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := WriteToOutbox(tx, pizzaOrder); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx.Commit()
	log.Println("Pizza order created successfully:", pizzaOrder)
	c.JSON(http.StatusOK, pizzaOrder)
}

func WriteToOutbox(tx *gorm.DB, pizzaOrder PizzaOrder) error {
	payload, err := json.Marshal(pizzaOrder)
	if err != nil {
		return err
	}

	outboxEntry := PizzaOrderOutbox{
		PizzaOrderID: pizzaOrder.ID,
		Status:       Pending,
		EventType:    PizzaOrderCreated,
		Payload:      string(payload),
	}

	return tx.Create(&outboxEntry).Error
}
