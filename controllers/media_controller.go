package controllers

import (
	"github.com/RodolfoBonis/upload_service/dtos"
	"github.com/RodolfoBonis/upload_service/models"
	"github.com/RodolfoBonis/upload_service/services"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"strings"
)

var mediaService = services.NewMediaUpload()

type MediaController interface {
	FileUpload(c echo.Context) error
	GetMedia(c echo.Context) error
}

type media struct{}

func NewMediaController() MediaController {
	return &media{}
}

func (*media) GetMedia(c echo.Context) error {
	bucketName := c.Param("bucket")
	mediaName := c.Param("media")

	if bucketName == "" || mediaName == "" {
		return c.NoContent(http.StatusNotFound)
	}

	mediaBuff, err := mediaService.GetMedia(c.Request().Context(), bucketName, mediaName)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	defer mediaBuff.Close()

	contentType := "application/octet-stream"

	extensionToContentType := map[string]string{
		"jpg":  "image/jpeg",
		"jpeg": "image/jpeg",
		"png":  "image/png",
		"mp4":  "video/mp4",
	}

	extension := strings.Split(mediaName, ".")[1]

	if ct, found := extensionToContentType[extension]; found {
		contentType = ct
	}

	c.Response().Header().Set("Content-Type", contentType)

	_, err = io.Copy(c.Response().Writer, mediaBuff)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (*media) FileUpload(c echo.Context) error {
	formHeader, err := c.FormFile("file")
	bucketName := c.FormValue("bucket")

	if bucketName == "" {
		return c.JSON(http.StatusBadRequest,
			dtos.MediaDto{
				StatusCode: http.StatusBadRequest,
				Message:    "Error",
				Data:       "Set a valid bucket name",
			})
	}

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

	uploadUrl, err := mediaService.FileUpload(models.FileModel{
		File: formFile,
		Name: formHeader.Filename,
		Size: formHeader.Size,
	},
		bucketName,
	)

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
