package controllers

import (
	"net/http"

	d "github.com/ch0ppy35/beer-docs/internal/database"
	"github.com/ch0ppy35/beer-docs/internal/models"
	"github.com/gin-gonic/gin"
)

type BreweryController struct{}

type BreweryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewBreweryController(e *gin.Engine) {
	b := BreweryController{}
	v1 := e.Group("/api/v1/breweries")
	{
		v1.POST("/", b.CreateBrewery)
		v1.GET("/", b.GetBreweries)
		v1.GET("/:id", b.GetSingleBrewery)
		v1.PUT("/:id", b.UpdateBrewery)
		v1.DELETE("/:id", b.DeleteBrewery)
	}
}

// POST /breweries
// Create a brewery
func (b *BreweryController) CreateBrewery(c *gin.Context) {
	var input BreweryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ensure input is correct! " + err.Error()})
		return
	}

	brewery := models.BreweryModel{Name: input.Name}
	d.Database.Create(&brewery)

	c.JSON(http.StatusOK, gin.H{"brewery": BreweryResponse{
		ID:   brewery.ID,
		Name: brewery.Name,
	}})
}

// GET /breweries
// Get all Breweries
func (b *BreweryController) GetBreweries(c *gin.Context) {
	var breweries []models.BreweryModel
	d.Database.Find(&breweries)
	// Create a new slice of BreweryResponse structs
	var breweryResponses []BreweryResponse
	// Map values from Brewery to BreweryResponse
	for _, brewery := range breweries {
		breweryResponses = append(breweryResponses, BreweryResponse{
			ID:   brewery.ID,
			Name: brewery.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{"breweries": breweryResponses})
}

// GET /breweries/:id
// Find a brewery
func (b *BreweryController) GetSingleBrewery(c *gin.Context) {
	var brewery models.BreweryModel

	if err := d.Database.Where("id = ?", c.Param("id")).First(&brewery).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	// Map values from Brewery to BreweryResponse
	breweryResponse := BreweryResponse{
		ID:   brewery.ID,
		Name: brewery.Name,
	}
	c.JSON(http.StatusOK, gin.H{"brewery": breweryResponse})
}

// PATCH /breweries/:id
// Update a brewery
func (b *BreweryController) UpdateBrewery(c *gin.Context) {
	var brewery models.BreweryModel
	if err := d.Database.Where("id = ?", c.Param("id")).First(&brewery).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input BreweryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ensure input is correct! " + err.Error()})
		return
	}

	d.Database.Model(&brewery).Updates(models.BreweryModel{Name: input.Name})
	c.JSON(http.StatusOK, gin.H{"data": brewery})
}

// DELETE /breweries/:id
// Delete a brewery, but only delete if there are no beers associated
func (b *BreweryController) DeleteBrewery(c *gin.Context) {
	var brewery models.BreweryModel
	if err := d.Database.Where("id = ?", c.Param("id")).Preload("Beers").First(&brewery).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	if len(brewery.Beers) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot delete brewery with associated beers!"})
		return
	}

	d.Database.Delete(&brewery)
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
