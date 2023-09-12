package services

import (
	"context"
	"github.com/RodolfoBonis/upload_service/helper"
	"github.com/RodolfoBonis/upload_service/models"
	"github.com/go-playground/validator/v10"
	"github.com/minio/minio-go/v7"
)

var (
	validate = validator.New()
)

type MediaUpload interface {
	FileUpload(file models.FileModel, bucketName string) (string, error)
	GetMedia(ctx context.Context, bucketName, mediaName string) (*minio.Object, error)
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

func (*media) GetMedia(ctx context.Context, bucketName, mediaName string) (*minio.Object, error) {

	mediaBuff, err := helper.GetMediaHelper(ctx, bucketName, mediaName)
	if err != nil {
		return nil, err
	}
	return mediaBuff, nil
}
