package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	host := os.Getenv("DBASE_HOST")
	password := os.Getenv("DBASE_PASSWORD")
	dbname := os.Getenv("DBASE_NAME")
	user := os.Getenv("DBASE_USER")
	var dbasePort string = os.Getenv("DBASE_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, dbasePort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to Connect to Database")
	}
}
