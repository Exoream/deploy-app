package storage

import (
	"context"
	"encoding/base64"
	"io"
	"mime/multipart"
	"os"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func UploadImageForPickup(image *multipart.FileHeader) (string, error) {
	ctx := context.Background()

	// Decode Google Cloud credentials from Base64
	encodedCredentials := os.Getenv("GOOGLE_CLOUD_CREDENTIALS_PATH")
	decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
	if err != nil {
		logrus.Error("Failed to decode Google Cloud credentials:", err)
		return "", err
	}

	client, err := storage.NewClient(ctx, option.WithCredentialsJSON(decodedCredentials))
	if err != nil {
		logrus.Error("Failed to create GCP client:", err)
		return "", err
	}
	defer client.Close()

	bucketName := "garvice"
	imagePath := "file_pickup/" + uuid.New().String() + ".jpg"

	wc := client.Bucket(bucketName).Object(imagePath).NewWriter(ctx)
	defer wc.Close()

	file, err := image.Open()
	if err != nil {
		logrus.Error("Failed to open image file:", err)
		return "", err
	}

	if _, err := io.Copy(wc, file); err != nil {
		logrus.Error("Failed to copy image data:", err)
		return "", err
	}

	imageURL := "https://storage.googleapis.com/" + bucketName + "/" + imagePath

	return imageURL, nil
}