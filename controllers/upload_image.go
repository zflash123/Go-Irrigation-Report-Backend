package controllers

import(
	"fmt"
	"strings"
	"regexp"
	"crypto/rand"
	"encoding/hex"
	"encoding/base64"
	"os"
	"go-irrigation-report-backend/models"
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

func UploadImage(image string) (fileUrl string, err error) {
	parts := strings.Split(image, ";")
	mimePart := strings.Split(parts[0], ":")
	mime := mimePart[1]
	imageExtension := strings.Split(mime, "/")[1]
	if(imageExtension=="go" || imageExtension=="svg"){
		return "", fmt.Errorf("The extension is prohibited.")
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
	defer destination.Close()

	fmt.Fprintf(destination, "%s", strDecodedImg)
	fileUrl, errUploadToFB := UploadToFirebase(imagePath, imageName)
	if errUploadToFB!=nil {
		return "", errUploadToFB
	}
	fileStat, _ := destination.Stat()
	fileSize := fileStat.Size()
	var uploadDump = models.UploadDump{
		Filename: imageName,
		FileType: imageExtension,
		Size: uint32(fileSize),
		Folder: "root",
		FileUrl: fileUrl,
	}
	models.Db.Create(&uploadDump)
	return fileUrl, nil
}