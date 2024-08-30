package utils

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"

	//_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

type ImgUtils struct {
	path string
}

func NewImgUtils(path string) *ImgUtils {
	return &ImgUtils{path: path}
}

type ImgUtilsService interface {
	LoadImage() (image.Image, error)
	ConvertToTensor(img image.Image) [][]color.Color
	CreateNewImage(name string, pixels [][]color.Color) error
}

func (i *ImgUtils) LoadImage() (image.Image, error) {
	f, err := os.Open(i.path)
	if err != nil {

		return nil, err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {

		return nil, err
	}
	fmt.Println(fi.Name())

	img, format, err := image.Decode(f)
	if err != nil {

		return nil, err
	}

	log.Println("Format is", format)
	if format != "jpeg" {

		return nil, errors.New("image format is not jpeg")
	}
	return img, nil
}

func (i *ImgUtils) ConvertToTensor(img image.Image) [][]color.Color {
	size := img.Bounds().Size()

	log.Println("Size is ", size)

	var pixels [][]color.Color

	for i := 0; i < size.X; i++ {
		var y []color.Color

		for j := 0; j < size.Y; j++ {

			y = append(y, img.At(i, j))
		}
		pixels = append(pixels, y)
	}
	return pixels
}

func (i *ImgUtils) CreateNewImage(name string, pixels [][]color.Color) error {
	rect := image.Rect(0, 0, len(pixels), len(pixels[0]))
	nImg := image.NewRGBA(rect)

	for x := 0; x < len(pixels); x++ {
		for y := 0; y < len(pixels[0]); y++ {
			q := pixels[x]
			if q == nil {
				continue
			}
			p := pixels[x][y]
			if p == nil {
				continue
			}
			original, ok := color.RGBAModel.Convert(p).(color.RGBA)
			if ok {
				nImg.Set(x, y, original)
			}
		}
	}
	dir := "./images/"

	fileName := fmt.Sprintf("%s%s.jpg", dir, name)

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Creating file:", err)
		return err
	}

	// Encode the image and write it to the file.
	err = jpeg.Encode(file, nImg, nil)
	if err != nil {
		fmt.Println("Error encoding image:", err)
		return err
	}

	return nil
}
