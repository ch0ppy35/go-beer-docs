package cmd

import (
	"os"
	"os/signal"
	"syscall"

	l "github.com/sirupsen/logrus"

	docs "github.com/ch0ppy35/beer-docs/docs"
	"github.com/ch0ppy35/beer-docs/internal/controllers"
	"github.com/ch0ppy35/beer-docs/internal/middleware"
	"github.com/ch0ppy35/beer-docs/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	// initialize logger
	utils.SetupJsonLogger()

	// create serve command
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Start server",
		Long:  `This starts the http server`,
		Run: func(cmd *cobra.Command, args []string) {
			StartServer()
		},
	}

	RootCmd.AddCommand(serveCmd)
}

func StartServer() {
	gin.SetMode(gin.ReleaseMode)
	p := utils.GetEnv("HTTP_PORT", "8080")
	r := gin.New()
	// Setup middleware
	r.Use(gin.Recovery())
	r.Use(middleware.JSONLogMiddleware())
	r.Use(middleware.RequestID(middleware.RequestIDOptions{AllowSetting: false}))
	r.Use(middleware.CORS(middleware.CORSOptions{Origin: "*"}))
	// gin-swagger middleware
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// Start db connection and migrations
	LoadAndMigrateDatabase()
	// Register controllers
	controllers.NewBreweryController(r)
	controllers.NewBeerController(r)
	controllers.NewHealthzController(r)
	// Start the server
	go func() {
		if err := r.Run(":" + p); err != nil {
			l.Fatalf("Server failed to start: %v", err)
		}
	}()
	// Block the main thread of execution until a signal is received
	// (e.g. SIGINT or SIGTERM)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
