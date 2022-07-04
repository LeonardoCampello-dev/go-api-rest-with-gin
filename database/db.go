package database

import (
	"log"

	"github.com/LeonardoCampello-dev/go-api-rest-with-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectWithDatabase() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Panic("Error connecting to database")
	}

	DB.AutoMigrate(&models.Student{})
}
