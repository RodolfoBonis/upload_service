package routes

import (
	"github.com/RodolfoBonis/upload_service/handlers"
	"github.com/RodolfoBonis/upload_service/middlewares"
	"github.com/labstack/echo/v4"
)

var uploadHandler = handlers.NewUploadHandler()
var healthHandler = handlers.NewHealthHandler()

type RootRoute interface {
	StartRoute(route *echo.Echo)
}

type routing struct{}

func NewRootRoute() RootRoute {
	return &routing{}
}

func (_ routing) StartRoute(route *echo.Echo) {

	route.GET("/health", healthHandler.GetHealth)
	route.GET("/:bucket/:media", uploadHandler.GetMedia)
	group := route.Group("/v1", middlewares.KeyGuardian)
	group.POST("/upload", uploadHandler.SaveImage)
}
