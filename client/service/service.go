package service

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

const pizzaURL = "http://localhost:8080/pizza"

var flavors = []string{"Margherita", "Pepperoni", "Hawaiian", "Veggie"}
var sizes = []string{"Small", "Medium", "Large"}
var addresses = []string{"123 Pizza St", "456 Pasta Ave", "789 Burger Blvd"}
var userNames = []string{"John Doe", "Jane Smith", "Alice Johnson"}

func CreatePizzaOrder() {
	// Simulate creating a pizza order
	pizzaOrder := PizzaOrder{
		Flavor:   flavors[rand.Intn(len(flavors))],
		Size:     sizes[rand.Intn(len(sizes))],
		Quantity: rand.Intn(5) + 1,      // Random quantity between 1 and 5
		Price:    rand.Float64() * 20.0, // Random price between 0 and 20
		Address:  addresses[rand.Intn(len(addresses))],
		UserName: userNames[rand.Intn(len(userNames))],
	}

	// Call the API to create the pizza order
	payload, err := json.Marshal(pizzaOrder)
	if err != nil {
		log.Fatal("Error marshalling pizza order:", err)
	}
	response, err := http.Post(pizzaURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal("Error posting pizza order:", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Println("Failed to create pizza order:", response.Status)
		return
	}

	log.Println("Pizza order created successfully!")
}
