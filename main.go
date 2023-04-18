package main

import (
	"fmt"
	config "github.com/RodolfoBonis/upload_service/configs"
	"github.com/RodolfoBonis/upload_service/routes"
	"github.com/labstack/echo/v4"
)

var rootRoute = routes.NewRootRoute()

func main() {
	config.LoadEnvVars()

	e := echo.New()

	rootRoute.StartRoute(e)

	port := config.EnvPortApplication()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))

}
