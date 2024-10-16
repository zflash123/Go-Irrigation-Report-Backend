package controllers

import (
	"fmt"
	"context"
	"os"
	"github.com/spf13/viper"
)

func UploadToFirebase(imagePath string, objectName string) error {
	viperEnvConfig()
	bucket := fmt.Sprintf("%v", viper.Get("GCS_BUCKET"))
	ctx := context.Background()
	wd, _ := os.Getwd()
	credentialsPath := fmt.Sprintf("%v/firebase-storage-credentials.json", wd)
}