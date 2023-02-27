package database

import (
	"fmt"

	"github.com/ch0ppy35/beer-docs/internal/utils"
	l "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB

func Connect() {
	var err error
	host := utils.GetEnv("DB_HOST", "127.0.0.1")
	username := utils.GetEnv("DB_USER", "postgres")
	password := utils.GetEnv("DB_PASSWORD", "mysecretpassword")
	databaseName := utils.GetEnv("DB_NAME", "svc_beer")
	port := utils.GetEnv("DB_PORT", "5432")

	connectionstring := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, username, password, databaseName, port)
	Database, err = gorm.Open(postgres.Open(connectionstring), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		panic(err)
	} else {
		l.Info("Connected to the database...")
	}
}
