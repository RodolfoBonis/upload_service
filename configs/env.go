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

func EnvCloudName() string {
	return GetEnv("CLOUDINARY_CLOUD_NAME", "")
}

func EnvCloudAPIKey() string {
	return GetEnv("CLOUDINARY_API_KEY", "")
}

func EnvCloudAPISecret() string {
	return GetEnv("CLOUDINARY_API_SECRET", "")
}

func EnvCloudUploadFolder() string {
	return GetEnv("CLOUDINARY_UPLOAD_FOLDER", "")
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
