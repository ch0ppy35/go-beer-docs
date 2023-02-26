package controllers

import (
	"net/http"

	d "github.com/ch0ppy35/beer-docs/database"
	"github.com/ch0ppy35/beer-docs/models"
	"github.com/gin-gonic/gin"
)

type CreateBeerInput struct {
	BeerName string `json:"beername" binding:"required"`
	Brewery  string `json:"brewery" binding:"required"`
}

// GET /beers/:id
// Find a beer
func GetSingleBeer(c *gin.Context) {
	var beer models.BeerModel

	if err := d.Database.Where("id = ?", c.Param("id")).First(&beer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": beer})
}

// GET /beers
// Get all Beers
func GetBeers(c *gin.Context) {
	var Beers []models.BeerModel
	d.Database.Find(&Beers)
	c.JSON(http.StatusOK, gin.H{"Response": Beers})
}

// POST /beers
// Create new beer
func CreateBeer(c *gin.Context) {
	// Validate input
	var input CreateBeerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong — " + err.Error()})
		return
	}

	beer := models.BeerModel{BeerName: input.BeerName, Brewery: input.Brewery}
	d.Database.Create(&beer)
	c.JSON(http.StatusOK, gin.H{"Response": beer})
}
