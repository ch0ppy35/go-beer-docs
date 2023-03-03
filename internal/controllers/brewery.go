package controllers

import (
	"net/http"

	d "github.com/ch0ppy35/beer-docs/internal/database"
	"github.com/ch0ppy35/beer-docs/internal/models"
	"github.com/gin-gonic/gin"
)

type BreweryController struct{}

type BreweryInput struct {
	Name string `json:"name" binding:"required"`
}

type BreweryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type DeletedBreweryResponse struct {
	Deleted bool `json:"deleted"`
}

type BreweryErrorResponse struct {
	Error string `string:"error"`
}

func NewBreweryController(e *gin.Engine) {
	b := BreweryController{}
	v1 := e.Group("/api/v1/breweries")
	{
		v1.POST("", b.CreateBrewery)
		v1.GET("", b.GetBreweries)
		v1.GET("/:id", b.GetSingleBrewery)
		v1.PATCH("/:id", b.UpdateBrewery)
		v1.DELETE("/:id", b.DeleteBrewery)
	}
}

// CreateBrewery godoc
// @Summary Create a brewery
// @Tags Breweries
// @Accept json
// @Produce json
// @Param brewery body BreweryInput true "Brewery input payload"
// @Success 200 {object} BreweryResponse
// @Failure 400 {object} BreweryErrorResponse
// @Router /breweries [post]
func (b *BreweryController) CreateBrewery(c *gin.Context) {
	var input BreweryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, BreweryErrorResponse{Error: "Ensure input is correct! " + err.Error()})
		return
	}

	brewery := models.BreweryModel{Name: input.Name}
	d.Database.Create(&brewery)

	c.JSON(http.StatusOK, BreweryResponse{
		ID:   brewery.ID,
		Name: brewery.Name,
	})
}

// GetBreweries godoc
// @Summary Get a list of all breweries
// @Description Get a list of all breweries
// @Tags Breweries
// @Accept json
// @Produce json
// @Success 200 {array} BreweryResponse
// @Failure 404 {object} BreweryErrorResponse
// @Router /breweries [get]
func (b *BreweryController) GetBreweries(c *gin.Context) {
	var breweries []models.BreweryModel
	d.Database.Find(&breweries)
	if len(breweries) == 0 {
		c.JSON(http.StatusNotFound, BreweryErrorResponse{Error: "No records found!"})
		return
	}
	// Create a new slice of BreweryResponse structs
	var breweryResponses []BreweryResponse
	// Map values from Brewery to BreweryResponse
	for _, brewery := range breweries {
		breweryResponses = append(breweryResponses, BreweryResponse{
			ID:   brewery.ID,
			Name: brewery.Name,
		})
	}

	c.JSON(http.StatusOK, breweryResponses)
}

// GetSingleBrewery godoc
// @Summary Get a single brewery by ID
// @Tags Breweries
// @Param id path int true "Brewery ID"
// @Accept json
// @Produce json
// @Success 200 {object} BreweryResponse
// @Failure 404 {object} BreweryErrorResponse
// @Router /breweries/{id} [get]
func (b *BreweryController) GetSingleBrewery(c *gin.Context) {
	var brewery models.BreweryModel

	if err := d.Database.Where("id = ?", c.Param("id")).First(&brewery).Error; err != nil {
		c.JSON(http.StatusNotFound, BreweryErrorResponse{Error: "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, BreweryResponse{
		ID:   brewery.ID,
		Name: brewery.Name,
	})
}

// UpdateBrewery godoc
// @Summary Update a brewery
// @Description Update a brewery by id
// @Tags Breweries
// @Accept json
// @Produce json
// @Param id path int true "Brewery ID"
// @Param brewery body BreweryInput true "Brewery Payload"
// @Success 200 {object} BreweryResponse
// @Failure 400 {object} BreweryErrorResponse
// @Failure 404 {object} BreweryErrorResponse
// @Router /breweries/{id} [patch]
func (b *BreweryController) UpdateBrewery(c *gin.Context) {
	var brewery models.BreweryModel
	if err := d.Database.Where("id = ?", c.Param("id")).First(&brewery).Error; err != nil {
		c.JSON(http.StatusNotFound, BreweryErrorResponse{Error: "Record not found!"})
		return
	}
	var input BreweryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, BreweryErrorResponse{Error: "Ensure input is correct! " + err.Error()})
		return
	}

	d.Database.Model(&brewery).Updates(models.BreweryModel{Name: input.Name})
	c.JSON(http.StatusOK, BreweryResponse{Name: input.Name})
}

// DeleteBrewery godoc
// @Summary Delete a brewery
// @Description Delete a brewery only if there are no beers associated with it
// @Tags Breweries
// @Produce json
// @Param id path int true "Brewery ID"
// @Success 200 {object} DeletedBreweryResponse
// @Failure 400 {object} BreweryErrorResponse
// @Failure 404 {object} BreweryErrorResponse
// @Router /breweries/{id} [delete]
func (b *BreweryController) DeleteBrewery(c *gin.Context) {
	var brewery models.BreweryModel
	if err := d.Database.Where("id = ?", c.Param("id")).Preload("Beers").First(&brewery).Error; err != nil {
		c.JSON(http.StatusNotFound, BreweryErrorResponse{Error: "Record not found!"})
		return
	}
	if len(brewery.Beers) > 0 {
		c.JSON(http.StatusBadRequest, BreweryErrorResponse{Error: "Cannot delete brewery with associated beers!"})
		return
	}

	d.Database.Delete(&brewery)
	c.JSON(http.StatusOK, DeletedBreweryResponse{Deleted: true})
}
