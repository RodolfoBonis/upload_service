package services

import (
	"github.com/RodolfoBonis/upload_service/helper"
	"github.com/RodolfoBonis/upload_service/models"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

type mediaUpload interface {
	FileUpload(file models.FileModel) (string, error)
	RemoteUpload(url models.UrlModel) (string, error)
}

type media struct{}

func NewMediaUpload() mediaUpload {
	return &media{}
}

func (*media) FileUpload(file models.FileModel) (string, error) {
	err := validate.Struct(file)

	if err != nil {
		return "", err
	}

	uploadUrl, err := helper.ImageUploadHelper(file.File)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}

func (*media) RemoteUpload(url models.UrlModel) (string, error) {
	err := validate.Struct(url)

	if err != nil {
		return "", err
	}

	uploadUrl, errUrl := helper.ImageUploadHelper(url.Url)

	if errUrl != nil {
		return "", errUrl
	}
	return uploadUrl, nil
}
