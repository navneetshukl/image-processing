package service

import (
	"image-processing/helpers"
	"image-processing/utils"
	"log"
)

type Image struct {
	Utils   utils.ImgUtilsService
	Helpers helpers.HelperService
}

func NewImageService(utils utils.ImgUtilsService, helpers helpers.HelperService) *Image {
	return &Image{Utils: utils, Helpers: helpers}
}

type ImageService interface {
	Invert() error
	RotateBy90D() error
	RotateBy270D() error
}

func (i *Image) Invert() error {

	img, err := i.Utils.LoadImage()
	if err != nil {
		log.Println("error in loading the image ", err)
		return err
	}

	pixels := i.Utils.ConvertToTensor(img)

	invertedPixels := i.Helpers.Invert(pixels)

	err = i.Utils.CreateNewImage("invert", invertedPixels)
	if err != nil {
		log.Println("error in creating new inverted image ", err)
		return err
	}

	return nil
}

func (i *Image) RotateBy90D() error {
	img, err := i.Utils.LoadImage()
	if err != nil {
		log.Println("error in loading the image ", err)
		return err
	}

	pixels := i.Utils.ConvertToTensor(img)

	rotatedPixel := i.Helpers.RotateBy90(pixels)
	err = i.Utils.CreateNewImage("90degree", rotatedPixel)
	if err != nil {
		log.Println("error in creating new 90 degree rotated image ", err)
		return err
	}

	return nil
}

func (i *Image) RotateBy270D() error {
	img, err := i.Utils.LoadImage()
	if err != nil {
		log.Println("error in loading the image ", err)
		return err
	}

	pixels := i.Utils.ConvertToTensor(img)
	rotatedPixels1 := i.Helpers.RotateBy90(pixels)
	rotatedPixel2 := i.Helpers.RotateBy90(rotatedPixels1)
	rotatedPixel3:=i.Helpers.RotateBy90(rotatedPixel2)

	err = i.Utils.CreateNewImage("270degree", rotatedPixel3)
	if err != nil {
		log.Println("error in creating new 270 degree rotated image ", err)
		return err
	}
	return nil
}
