package models

import "gorm.io/gorm"

type BeerModel struct {
	gorm.Model
	BeerName string `json:"beername"`
	Brewery  string `json:"brewery"`
}
