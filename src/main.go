package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/rakyll/statik/fs"

	Petstore "github.com/abhishekb91/petstore-openapi3/src/petstore"
	_ "github.com/abhishekb91/petstore-openapi3/src/statik" //swagger-ui loaded via statik
)

func main() {
	e := echo.New()
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	// Log all requests
	e.Use(echomiddleware.Logger())

	// Serve swagger-ui contents over HTTP.
	sh := http.StripPrefix("/docs/", http.FileServer(statikFS))
	eh := echo.WrapHandler(sh)
	e.GET("/docs/*", eh)

	// Get api specification to parse in swagger-ui.
	e.GET("docs/openapi.json", func(ctx echo.Context) error {
		spec, err := Petstore.GetSwagger()

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		ctx.JSON(200, spec)
		return nil
	})

	e.Logger.Fatal(e.Start(":8080"))
}
