package cmd

import (
	d "github.com/ch0ppy35/beer-docs/internal/database"
	"github.com/ch0ppy35/beer-docs/internal/models"
	l "github.com/sirupsen/logrus"
)

func LoadAndMigrateDatabase() {
	d.Connect()
	l.Info("Running database migrations...")
	d.Database.AutoMigrate(&models.BeerModel{}, &models.BreweryModel{})
	l.Info("Database migrations complete!")
}
