package main

import (
	"image-processing/helpers"
	"image-processing/service"
)

func main() {
	imgHelperService := helpers.NewImgHelper("./images/image1.jpg")
	imageService := service.NewImageService(imgHelperService)

	imageService.Invert()
}
