package main

import (
	"github.com/RodolfoBonis/upload_service/controllers"
	"github.com/labstack/echo/v4"
)

var mediaController = controllers.NewMediaController()

func main() {
	e := echo.New()

	e.POST("/upload", func(c echo.Context) error {
		_, err := c.FormFile("file")

		if err != nil {
			return mediaController.RemoteUpload(c)
		} else {
			return mediaController.FileUpload(c)
		}
	})

	e.Logger.Fatal(e.Start(":6000"))

}
