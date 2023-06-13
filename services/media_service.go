package services

import (
	"github.com/RodolfoBonis/upload_service/helper"
	"github.com/RodolfoBonis/upload_service/models"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

type MediaUpload interface {
	FileUpload(file models.FileModel, bucketName string) (string, error)
}

type media struct{}

func NewMediaUpload() MediaUpload {
	return &media{}
}

func (*media) FileUpload(file models.FileModel, bucketName string) (string, error) {
	err := validate.Struct(file)

	if err != nil {
		return "", err
	}

	uploadUrl, err := helper.ImageUploadHelper(file, bucketName)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}
