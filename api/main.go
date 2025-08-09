package main

import (
	"log"
	"net/http"
	"outbox/db"
	"outbox/service"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

	db.Connect()

	err := db.DB.AutoMigrate(&service.PizzaOrder{}, &service.PizzaOrderOutbox{})
	if err != nil {
		log.Fatal("Failed to migrate database")
	}

	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.POST("/pizza", service.CreatePizza)

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo": "bar", // user:foo password:bar
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
