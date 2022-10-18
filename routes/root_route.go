package routes

import (
	"github.com/RodolfoBonis/upload_service/handlers"
	"github.com/labstack/echo/v4"
)

var uploadHandler = handlers.NewUploadHandler()

type RootRoute interface {
	StartRoute(route *echo.Echo)
}

type routing struct{}

func NewRootRoute() RootRoute {
	return &routing{}
}

func (_ routing) StartRoute(route *echo.Echo) {
	group := route.Group("/v1")

	group.POST("/upload", uploadHandler.SaveImage)

}
