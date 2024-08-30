package helpers

import (
	"image/color"
)

type Helper struct {
}

func NewHelper() *Helper {
	return &Helper{}
}

type HelperService interface {
	RotateBy90(pixels [][]color.Color) [][]color.Color
	Invert(pixels [][]color.Color) [][]color.Color
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
