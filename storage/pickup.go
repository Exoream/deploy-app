package storage

import (
	"context"
	"encoding/base64"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func UploadImageForPickup(image *multipart.FileHeader) (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		logrus.Error("Config : Cannot load config file,", err.Error())
	}

	ctx := context.Background()

	// keys credentials google cloud
	credentialBase64 := os.Getenv("GOOGLE_CLOUD_CREDENTIALS_BASE64")
	credentialBytes, err := base64.StdEncoding.DecodeString(credentialBase64)
	if err != nil {
		return "", err
	}

	credentialFile := "keys.json"
	err = ioutil.WriteFile(credentialFile, credentialBytes, 0644)
	if err != nil {
		return "", err
	}
	defer os.Remove(credentialFile)

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFile))
	if err != nil {
		return "", err
	}
	defer client.Close()

	bucketName := "garvice"
	imagePath := "file_pickup/" + uuid.New().String() + ".jpg"

	wc := client.Bucket(bucketName).Object(imagePath).NewWriter(ctx)
	defer wc.Close()
	file, err := image.Open()
	if err != nil {
		return "", err
	}
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}

	imageURL := "https://storage.googleapis.com/" + bucketName + "/" + imagePath

	return imageURL, nil
}
