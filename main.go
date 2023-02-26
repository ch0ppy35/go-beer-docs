package main

import (
	"log"

	"github.com/ch0ppy35/beer-docs/controllers"
	"github.com/ch0ppy35/beer-docs/database"
	"github.com/ch0ppy35/beer-docs/models"
	"github.com/gin-gonic/gin"
)

func main() {
	loadAndMigrateDatabase()
	r := gin.Default()
	v1 := r.Group("/api/v1/beers")
	{
		v1.POST("/", controllers.CreateBeer)
		v1.GET("/", controllers.GetBeers)
		v1.GET("/:id", controllers.GetSingleBeer)
		// v1.PUT("/:id", updateBeer)
		// v1.DELETE("/:id", deleteBeer)
	}
	r.Run()
}

func loadAndMigrateDatabase() {
	database.Connect()
	log.Println("Connected to the database...")
	log.Println("Running database migrations...")
	database.Database.AutoMigrate(&models.BeerModel{})
	log.Printf("Database migrations complete!")
}
