package database

import (
	"fmt"
	"log"

	"github.com/ch0ppy35/beer-docs/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	var err error
	host := utils.GetEnv("DB_HOST", "127.0.0.1")
	username := utils.GetEnv("DB_USER", "postgres")
	password := utils.GetEnv("DB_PASSWORD", "mysecretpassword")
	databaseName := utils.GetEnv("DB_NAME", "svc_beer")
	port := utils.GetEnv("DB_PORT", "5432")

	connectionstring := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, databaseName, port)
	Database, err = gorm.Open(postgres.Open(connectionstring), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		log.Println("Successfully connected to the database")
	}
}
