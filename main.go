package main

import (
	"bitbucket.org/bri_bootcamp/fp-patungan-backend-go/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	
	e := echo.New()
	routes.SetupRoutes(e)
	e.Logger.Fatal(e.Start("localhost:9090"))

}
