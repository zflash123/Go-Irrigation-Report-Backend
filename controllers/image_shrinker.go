package controllers

import "os/exec"

func ShrinkImage(imagePath string, toSize string) (err error){
	err = exec.Command("convert", imagePath, "-quality", "100", "-resize", toSize, imagePath).Run()
	return err
}
