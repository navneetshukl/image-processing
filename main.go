package main

import (
	"image-processing/helpers"
	"image-processing/service"
	"image-processing/utils"
)

func main() {
	imgUtilsService := utils.NewImgUtils("./images/image1.jpg")
	helperService := helpers.NewHelper()
	imageService := service.NewImageService(imgUtilsService, helperService)

	//imageService.Invert()
	imageService.RotateBy90D()
	imageService.RotateBy270D()
}
