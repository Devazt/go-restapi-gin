package cloudinary

import (
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
)

func CloudinarySetup() (*cloudinary.Cloudinary, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cloudName := os.Getenv("CLOUD_NAME")
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")

	cloud, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return nil, err
	}

	return cloud, nil
}
