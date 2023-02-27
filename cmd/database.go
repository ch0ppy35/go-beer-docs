package cmd

import (
	"github.com/ch0ppy35/beer-docs/internal/database"
	"github.com/ch0ppy35/beer-docs/internal/models"
	l "github.com/sirupsen/logrus"
)

func LoadAndMigrateDatabase() {
	database.Connect()
	l.Info("Running database migrations...")
	database.Database.AutoMigrate(&models.BeerModel{})
	l.Info("Database migrations complete!")
}
