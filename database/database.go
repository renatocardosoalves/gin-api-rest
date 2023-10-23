package database

import (
	"log"

	"github.com/renatocardosoalves/api-go-gin/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable TimeZone=UTC"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Error connecting to database")
	}
	DB.AutoMigrate(&model.Student{})
}
