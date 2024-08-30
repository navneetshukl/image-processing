package service

import (
	"image-processing/helpers"
	"log"
)

type Image struct {
	Helpers helpers.ImgHelperService
}

func NewImageService(helpers helpers.ImgHelperService) *Image {
	return &Image{helpers}
}

type ImageService interface {
	UpsideDown() error
}

func (i *Image) UpsideDown() error {

	img, err := i.Helpers.LoadImage()
	if err != nil {
		log.Println("error in loading the image ", err)
		return err
	}

	pixels := i.Helpers.ConvertToTensor(img)

	row := len(pixels)

	for i := 0; i < row; i++ {
		
		col := pixels[i]

		si := len(col)

		j, k := 0, si-1

		for j < k {
			col[j], col[k] = col[k], col[j]
			j++
			k--
		}
	}

	err = i.Helpers.CreateNewImage("invert", pixels)
	if err != nil {
		log.Println("error in creating new inverted image ", err)
		return err
	}

	return nil
}
