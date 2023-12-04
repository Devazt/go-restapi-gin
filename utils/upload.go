package utils

import (
	"context"
	"mime/multipart"

	"github.com/Devazt/go-restapi-gin/pkg/cloudinary"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadtoCloud(file multipart.File, filePath string) (string, error) {
	ctx := context.Background()
	cloud, err := cloudinary.CloudinarySetup()
	if err != nil {
		return "", err
	}

	uploadParams := uploader.UploadParams{
		PublicID: filePath,
	}

	result, err := cloud.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", err
	}
	imageUrl := result.SecureURL
	return imageUrl, nil
}
