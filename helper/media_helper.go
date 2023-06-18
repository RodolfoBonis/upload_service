package helper

import (
	"context"
	"fmt"
	config "github.com/RodolfoBonis/upload_service/configs"
	"github.com/RodolfoBonis/upload_service/models"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"os"
)

func ImageUploadHelper(file models.FileModel, bucketName string) (string, error) {
	ctx := context.Background()
	endpoint := config.EnvMinioServer()
	accessKeyID := config.EnvAccessID()
	secretAccessKey := config.EnvSecretKey()

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: true,
	})
	if err != nil {
		return "", err
	}

	_, err = minioClient.PutObject(
		ctx,
		bucketName,
		file.Name,
		file.File,
		file.Size,
		minio.PutObjectOptions{},
	)

	if err != nil {
		return "", err
	}

	os.Remove(file.Name)

	algorithm := "?X-Amz-Algorithm=AWS4-HMAC-SHA256"

	return fmt.Sprintf("https://%s/%s/%s%s", endpoint, bucketName, file.Name, algorithm), nil
}
