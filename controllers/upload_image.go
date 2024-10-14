package controllers

import(
	"strings"
)

func UploadImage(image string) error {
	parts := strings.Split(image, ";")
	mimePart := strings.Split(parts[0], ":")
	mime := mimePart[1]
}