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

type BeersResponse struct {
}

type DeletedBeerResponse struct {
	Deleted bool `json:"deleted"`
}

type BeerErrorResponse struct {
	Error string `string:"error"`
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

// CreateBeer godoc
// @Summary Create a new beer
// @Description Create a new beer
// @Tags Beers
// @Accept json
// @Produce json
// @Param input body BeerInput true "Beer input"
// @Success 200 {object} BeerResponse "Successful operation"
// @Failure 400 {object} BeerErrorResponse "Ensure input is correct!"
// @Router /beers [post]
func (b *BeerController) CreateBeer(c *gin.Context) {
	// Validate input
	var input BeerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, BeerErrorResponse{Error: "Ensure input is correct! " + err.Error()})
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

	c.JSON(http.StatusOK, beerResponse)
}

// GetBeers godoc
// @Summary Get a list of all beers
// @Description Get a list of all beers
// @Tags Beers
// @Accept json
// @Produce json
// @Success 200 {array} BeerResponse "Successful operation"
// @Failure 404 {object} BeerErrorResponse
// @Router /beers [get]
func (b *BeerController) GetBeers(c *gin.Context) {
	var beers []models.BeerModel
	d.Database.Preload("Brewery").Find(&beers)
	if len(beers) == 0 {
		c.JSON(http.StatusNotFound, BeerErrorResponse{Error: "No records found!"})
		return
	}

	// Create a new slice of BeerResponse structs
	var beerResponses []BeerResponse

	// Map values from BeerModel to BeerResponse
	for _, beer := range beers {
		beerResponses = append(beerResponses, BeerResponse{
			BeerName: beer.BeerName,
			Brewery:  BreweryInput{Name: beer.Brewery.Name},
		})
	}

	c.JSON(http.StatusOK, beerResponses)
}

// GetSingleBeer returns a single beer by ID
// @Summary Get a beer by ID
// @Tags Beers
// @Param id path int true "Beer ID"
// @Produce json
// @Success 200 {object} BeerResponse
// @Failure 404 {object} BeerErrorResponse
// @Router /beers/{id} [get]
func (b *BeerController) GetSingleBeer(c *gin.Context) {
	var beer models.BeerModel

	if err := d.Database.Preload("Brewery").Where("id = ?", c.Param("id")).First(&beer).Error; err != nil {
		c.JSON(http.StatusNotFound, BeerErrorResponse{Error: "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, BeerResponse{
		BeerName: beer.BeerName,
		Brewery:  BreweryInput{Name: beer.Brewery.Name},
	})
}

// UpdateBeer updates an existing beer record by ID
// @Summary Update a beer by ID
// @Tags Beers
// @Param id path int true "Beer ID"
// @Accept json
// @Produce json
// @Param beer body BeerInput true "Beer input payload"
// @Success 200 {object} BeerResponse
// @Failure 400 {object} BeerErrorResponse
// @Failure 404 {object} BeerErrorResponse
// @Router /beers/{id} [patch]
func (b *BeerController) UpdateBeer(c *gin.Context) {
	var beer models.BeerModel
	if err := d.Database.Where("id = ?", c.Param("id")).First(&beer).Error; err != nil {
		c.JSON(http.StatusNotFound, BeerErrorResponse{Error: "Record not found!"})
		return
	}

	var input BeerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, BeerErrorResponse{Error: "Ensure input is correct! " + err.Error()})
		return
	}

	// Look up or create the brewery
	var brewery models.BreweryModel
	d.Database.FirstOrCreate(&brewery, models.BreweryModel{Name: input.Brewery.Name})

	// Update the beer
	d.Database.Model(&beer).Updates(models.BeerModel{BeerName: input.BeerName, Brewery: brewery})

	// Map values from BeerModel to BeerResponse

	c.JSON(http.StatusOK, BeerResponse{
		BeerName: beer.BeerName,
		Brewery:  BreweryInput{Name: beer.Brewery.Name},
	})
}

// DeleteBeer deletes a beer by ID
// @Summary Delete a beer by ID
// @Tags Beers
// @Param id path int true "Beer ID"
// @Produce json
// @Success 200 {object} DeletedBeerResponse
// @Failure 400 {object} BeerErrorResponse
// @Router /beers/{id} [delete]
func (b *BeerController) DeleteBeer(c *gin.Context) {
	var beer models.BeerModel
	if err := d.Database.Where("id = ?", c.Param("id")).First(&beer).Error; err != nil {
		c.JSON(http.StatusBadRequest, BeerErrorResponse{Error: "Record not found!"})
		return
	}

	d.Database.Delete(&beer)

	c.JSON(http.StatusOK, DeletedBeerResponse{Deleted: true})
}
