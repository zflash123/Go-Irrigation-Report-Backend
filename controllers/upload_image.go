package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"go-irrigation-report-backend/models"
	"os"
	"regexp"
	"strings"
)

func GenerateCryptoID() string {
	bytes := make([]byte, 7)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	result := hex.EncodeToString(bytes)
	result = result[0:13]
	return result
}

func UploadImage(image string) (reportPhotoID string, err error) {
	// Get Image Extension from string of uploaded Image and then store it in variable imageExtension
	parts := strings.Split(image, ";")
	mimePart := strings.Split(parts[0], ":")
	imageExtension := (strings.Split(mimePart[1], "/"))[1]
	
	if(imageExtension=="go" || imageExtension=="svg"){
		return "", fmt.Errorf("the extension is prohibited")
	}
	regex, err := regexp.Compile(`(?i)data:image/[\w]+;base64,`)
	if err != nil {
		fmt.Println(err.Error())
	}
	res1 := regex.FindString(image)
	image = strings.Replace(image, res1, "", 1)
	uniqueId := GenerateCryptoID()
	imageName := fmt.Sprintf("%v.%v", uniqueId, imageExtension)
	wd, _ := os.Getwd()
	imagePath := fmt.Sprintf("%v/images/%v", wd, imageName)
	var decodedImg []byte
	decodedImg, _ = base64.StdEncoding.DecodeString(image)
	strDecodedImg := string(decodedImg)
	destination, _ := os.Create(imagePath)

	fmt.Fprintf(destination, "%s", strDecodedImg)
	fileUrl, errUploadToFB := UploadToFirebase(imagePath, imageName)
	if errUploadToFB!=nil {
		return "", errUploadToFB
	}
	fileStat, _ := destination.Stat()
	fileSize := fileStat.Size()
	destination.Close()
	var reportPhoto = models.ReportPhoto{
		Filename: imageName,
		FileType: imageExtension,
		Size: uint32(fileSize),
		FileUrl: fileUrl,
	}
	models.Db.Create(&reportPhoto)
	reportPhotoID = fmt.Sprintf("%s", reportPhoto.ID)
	err = os.Remove(imagePath)
	if err!=nil {
		fmt.Println("Err: ", err)
	}
	return reportPhotoID, nil
}