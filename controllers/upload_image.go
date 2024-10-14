package controllers

import(
	"fmt"
	"strings"
	"regexp"
)

func UploadImage(image string) error {
	parts := strings.Split(image, ";")
	mimePart := strings.Split(parts[0], ":")
	mime := mimePart[1]
	imageExtension := strings.Split(mime, "/")[1]
	if(imageExtension=="go" || imageExtension=="svg"){
		return fmt.Errorf("The extension is prohibited.")
	}
	regex, err := regexp.Compile(`(?i)data:image/[\w]+;base64,`)
}