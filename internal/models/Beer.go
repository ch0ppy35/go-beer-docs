package models

import (
	"gorm.io/gorm"
)

type BeerModel struct {
	gorm.Model
	BeerName  string `json:"beername"`
	BreweryID uint   `json:"brewery_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Brewery   BreweryModel
}
