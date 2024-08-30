package helpers

import (
	"image/color"
	"log"
)

type Helper struct {
}

func NewHelper() *Helper {
	return &Helper{}
}

type HelperService interface {
	RotateBy90(pixels [][]color.Color) [][]color.Color
	Invert(pixels [][]color.Color) [][]color.Color
	GreyScale(pixels [][]color.Color) [][]color.Color
}

func (h *Helper) RotateBy90(pixels [][]color.Color) [][]color.Color {
	row := len(pixels)
	col := len(pixels[0])

	// Create a new matrix for the result
	result := make([][]color.Color, col)
	for i := range result {
		result[i] = make([]color.Color, row)
	}

	// Transpose and reverse the rows
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			result[j][row-i-1] = pixels[i][j]
		}
	}

	return result
}

func (h *Helper) Invert(pixels [][]color.Color) [][]color.Color {
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
	return pixels

}

// 0.21 R + 0.72 G + 0.07 B
func (h *Helper) GreyScale(pixels [][]color.Color) [][]color.Color {
	row := len(pixels)
	col := len(pixels[0])
	result := make([][]color.Color, row)

	for i := range pixels {
		result[i] = make([]color.Color, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {

			originalColor, ok := color.RGBAModel.Convert(pixels[i][j]).(color.RGBA)
			if !ok {
				log.Println("type conversion went wrong")
				return nil
			}

			grey := uint8(float64(originalColor.R)*0.21 + float64(originalColor.G)*0.72 + float64(originalColor.B)*0.07)
			col := color.RGBA{
				grey,
				grey,
				grey,
				originalColor.A,
			}

			result[i][j] = col

		}
	}
	return result
}
