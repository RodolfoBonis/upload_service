package controllers

import (
	"net/http"

	"github.com/RodolfoBonis/upload_service/dtos"
	"github.com/RodolfoBonis/upload_service/models"
	"github.com/RodolfoBonis/upload_service/services"
	"github.com/labstack/echo/v4"
)

var mediaService = services.NewMediaUpload()

type MediaController interface {
	FileUpload(c echo.Context) error
	RemoteUpload(c echo.Context) error
}

type media struct{}

func NewMediaController() MediaController {
	return &media{}
}

func (*media) FileUpload(c echo.Context) error {
	formHeader, err := c.FormFile("file")

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error",
				Data:       "Select a File to Upload",
			})
	}

	formFile, err := formHeader.Open()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       err.Error(),
			})
	}

	uploadUrl, err := mediaService.FileUpload(models.FileModel{File: formFile})

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       err.Error(),
			})
	}

	return c.JSON(http.StatusOK,
		dtos.MediaDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       uploadUrl,
		})
}

func (*media) RemoteUpload(c echo.Context) error {
	var url models.UrlModel

	if err := c.Bind(&url); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			dtos.MediaDto{
				StatusCode: http.StatusBadRequest,
				Message:    "error",
				Data:       err.Error(),
			})
	}

	uploadUrl, err := mediaService.RemoteUpload(url)

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       "Error uploading file",
			})
	}

	return c.JSON(
		http.StatusOK,
		dtos.MediaDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       uploadUrl,
		})

}
