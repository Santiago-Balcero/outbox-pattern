package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const dsn = "host=localhost user=outbox_user password=outbox_pass dbname=outbox_db port=5432 sslmode=disable TimeZone=America/Bogota"

func Connect() {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
}
