package models

import "gorm.io/gorm"

type BreweryModel struct {
	gorm.Model
	Name  string
	Beers []BeerModel `gorm:"foreignKey:BreweryID"`
}
