package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	return defaultValue
}

func EnvMinioServer() string {
	return GetEnv("MINIO_SERVER", "")
}

func EnvAccessID() string {
	return GetEnv("ACCESS_ID", "")
}

func EnvSecretKey() string {
	return GetEnv("SECRET_KEY", "")
}

func EnvBucketName() string {
	return GetEnv("BUCKET_NAME", "")
}

func EnvPortApplication() string {
	return GetEnv("APP_PORT", ":8000")
}

func LoadEnvVars() {
	env := GetEnv("UPLOAD_ENV", "development")

	if env == "production" || env == "staging" {
		log.Println("Not using .env file in production or staging.")
		return
	}

	filename := ".env." + env

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		filename = ".env"
	}

	err := godotenv.Load(filename)
	if err != nil {
		log.Fatal(".env file not loaded")
	}
}
