//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=api --generate server,spec -o ./api/petstore-server.gen.go ./petstore.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=api --generate types -o ./api/petstore-type.gen.go ./petstore.yaml

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/abhishekb91/petstore-openapi3/src/models"
	_ "github.com/abhishekb91/petstore-openapi3/src/statik" //swagger-ui loaded via statik
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rakyll/statik/fs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/abhishekb91/petstore-openapi3/src/api"
	"github.com/abhishekb91/petstore-openapi3/src/controllers"
	"github.com/abhishekb91/petstore-openapi3/src/database"
)

func main() {
	e := echo.New()

	// Log all requests
	e.Use(middleware.Logger())

	//Connecting To DB
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("[main.main]: Failed to connect database, Error: %v", err)
	}
	dbObj, _ := db.DB()
	defer dbObj.Close()

	//Migrating Models
	db.AutoMigrate(&models.Pet{})

	//Registering API Docs
	serveApiDocs(e)

	//Registering Routes
	myApi := controllers.NewSvcController(database.NewDataAccessor(db))
	api.RegisterHandlers(e, myApi)

	// Start server
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("[main.main]: shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func serveApiDocs(e *echo.Echo) {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	// Serve swagger-ui contents over HTTP.
	sh := http.StripPrefix("/docs/", http.FileServer(statikFS))
	eh := echo.WrapHandler(sh)
	e.GET("/docs/*", eh)

	// Get api specification to parse in swagger-ui.
	e.GET("docs/openapi.json", func(ctx echo.Context) error {
		spec, err := api.GetSwagger()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		ctx.JSON(200, spec)
		return nil
	})
}
