package database

import (
	"fmt"
	"log"

	"github.com/Kozzen890/assignment3-016/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var(
	host     = "localhost"
	user     = "postgres"
	password = "postgres890"
	port     = 5432
	dbname   = "weatherapi"
)

func DbInit()(*gorm.DB, error){
	config := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", host, user, password, port, dbname)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("Error:", err)
	}
	db.Debug().AutoMigrate(models.WeatherData{})
	return db, nil
}
