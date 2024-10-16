package controllers

import (
	"fmt"
	"context"
	"github.com/spf13/viper"
)

func UploadToFirebase(imagePath string, objectName string) error {
	viperEnvConfig()
	bucket := fmt.Sprintf("%v", viper.Get("GCS_BUCKET"))
	ctx := context.Background()
}