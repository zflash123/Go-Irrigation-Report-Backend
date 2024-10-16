package controllers

import (
	"fmt"
	"context"
	"os"
	"github.com/spf13/viper"
	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func UploadToFirebase(imagePath string, objectName string) error {
	viperEnvConfig()
	bucket := fmt.Sprintf("%v", viper.Get("GCS_BUCKET"))
	ctx := context.Background()
	wd, _ := os.Getwd()
	credentialsPath := fmt.Sprintf("%v/firebase-storage-credentials.json", wd)
	opt := option.WithCredentialsFile(credentialsPath)
	client, err := storage.NewClient(ctx, opt)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()

	// Open local file.
	f, err := os.Open(imagePath)
}