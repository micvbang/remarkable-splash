package splash

import (
	"image"
	"image/color"

	"github.com/disintegration/imaging"
)

func Resize(img image.Image) (image.Image, error) {
	const (
		width  = 1404
		height = 1872
	)

	img = imaging.Resize(img, 1404, 0, imaging.Linear)

	container := imaging.New(
		width,
		height,
		color.RGBA{255, 255, 255, 255},
	)

	return imaging.PasteCenter(container, img), nil
}
