//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=api --generate server,spec -o ./api/petstore-server.gen.go ./petstore.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=api --generate types -o ./api/petstore-type.gen.go ./petstore.yaml

package main

import (
	"log"
	"net/http"

	_ "github.com/abhishekb91/petstore-openapi3/src/statik" //swagger-ui loaded via statik
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rakyll/statik/fs"

	"github.com/abhishekb91/petstore-openapi3/src/api"
	"github.com/abhishekb91/petstore-openapi3/src/controllers"
	"github.com/abhishekb91/petstore-openapi3/src/database"
)

func main() {
	e := echo.New()
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	// Log all requests
	e.Use(middleware.Logger())

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

	//Registering Routes
	myApi := controllers.NewSvcController(database.NewDataAccessor()) // This implements the pet store interface
	api.RegisterHandlers(e, myApi)

	e.Logger.Fatal(e.Start(":8080"))
}
