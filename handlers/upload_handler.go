package handlers

import (
	"github.com/RodolfoBonis/upload_service/controllers"
	"github.com/labstack/echo/v4"
	"net/http"
)

var mediaController = controllers.NewMediaController()

type UploadHandler interface {
	SaveImage(c echo.Context) error
}

type upload struct{}

func NewUploadHandler() UploadHandler {
	return &upload{}
}

func (_ upload) SaveImage(c echo.Context) error {
	_, err := c.FormFile("file")

	if err != nil {
		return c.String(http.StatusBadRequest, "Please, send a valid file")
	}

	return mediaController.FileUpload(c)
}
