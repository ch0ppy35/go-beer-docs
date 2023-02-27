package controllers

import (
	"net/http"

	d "github.com/ch0ppy35/beer-docs/internal/database"
	"github.com/ch0ppy35/beer-docs/internal/models"
	"github.com/gin-gonic/gin"
)

type BeerController struct{}

type BeerInput struct {
	BeerName string       `json:"beername" binding:"required"`
	Brewery  BreweryInput `json:"brewery" binding:"required"`
}

type BreweryInput struct {
	Name string `json:"name" binding:"required"`
}

type BeerResponse struct {
	BeerName string       `json:"name"`
	Brewery  BreweryInput `json:"brewery"`
}

func NewBeerController(e *gin.Engine) {
	b := BeerController{}
	v1 := e.Group("/api/v1/beers")
	{
		v1.POST("/", b.CreateBeer)
		v1.GET("/", b.GetBeers)
		v1.GET("/:id", b.GetSingleBeer)
		v1.PUT("/:id", b.UpdateBeer)
		v1.DELETE("/:id", b.DeleteBeer)
	}
}

// POST /beers
// Create new beer
func (b *BeerController) CreateBeer(c *gin.Context) {
	// Validate input
	var input BeerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ensure input is correct! " + err.Error()})
		return
	}

	// Look up or create the brewery
	var brewery models.BreweryModel
	d.Database.FirstOrCreate(&brewery, models.BreweryModel{Name: input.Brewery.Name})

	// Create the beer
	beer := models.BeerModel{BeerName: input.BeerName, Brewery: brewery}
	d.Database.Create(&beer)
	// Map values from BeerModel to BeerResponse
	beerResponse := BeerResponse{
		BeerName: beer.BeerName,
		Brewery:  BreweryInput{Name: beer.Brewery.Name},
	}

	c.JSON(http.StatusOK, gin.H{"Beer": beerResponse})
}

// GET /beers
// Get all Beers
func (b *BeerController) GetBeers(c *gin.Context) {
	var beers []models.BeerModel
	d.Database.Preload("Brewery").Find(&beers)

	// Create a new slice of BeerResponse structs
	var beerResponses []BeerResponse

	// Map values from BeerModel to BeerResponse
	for _, beer := range beers {
		beerResponses = append(beerResponses, BeerResponse{
			BeerName: beer.BeerName,
			Brewery:  BreweryInput{Name: beer.Brewery.Name},
		})
	}

	c.JSON(http.StatusOK, gin.H{"beers": beerResponses})
}

// GET /beers/:id
// Find a beer
func (b *BeerController) GetSingleBeer(c *gin.Context) {
	var beer models.BeerModel

	if err := d.Database.Preload("Brewery").Where("id = ?", c.Param("id")).First(&beer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	// Map values from BeerModel to BeerResponse
	beerResponse := BeerResponse{
		BeerName: beer.BeerName,
		Brewery:  BreweryInput{Name: beer.Brewery.Name},
	}

	c.JSON(http.StatusOK, gin.H{"beer": beerResponse})
}

// PATCH /beers/:id
// Update a beer
func (b *BeerController) UpdateBeer(c *gin.Context) {
	var beer models.BeerModel
	if err := d.Database.Where("id = ?", c.Param("id")).First(&beer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input BeerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ensure input is correct! " + err.Error()})
		return
	}

	// Look up or create the brewery
	var brewery models.BreweryModel
	d.Database.FirstOrCreate(&brewery, models.BreweryModel{Name: input.Brewery.Name})

	// Update the beer
	d.Database.Model(&beer).Updates(models.BeerModel{BeerName: input.BeerName, Brewery: brewery})

	c.JSON(http.StatusOK, gin.H{"data": beer})
}

// DELETE /beers/:id
// Delete a beer, but keep brewery in system
func (b *BeerController) DeleteBeer(c *gin.Context) {
	var beer models.BeerModel
	if err := d.Database.Where("id = ?", c.Param("id")).First(&beer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	d.Database.Delete(&beer)

	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
