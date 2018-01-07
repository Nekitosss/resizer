package resizer

import (
	"image"
	"github.com/nfnt/resize"
)


type ResizeRequest struct {

	URL    string `json:"url,omitempty"`

	Height int    `json:"height,omitempty"`

	Width  int    `json:"width,omitempty"`

}


func Resize(width int, height int, img *image.Image) *image.Image {
	original := *img
	originalWidth := original.Bounds().Max.X
	originalHeight := original.Bounds().Max.Y

	preferredWidth := uint(0)
	preferredHeight := uint(0)

	if originalWidth < originalHeight {
		preferredWidth = uint(width)
		preferredHeight = uint(float32(originalHeight * width) / float32(originalWidth))
	} else {
		preferredHeight = uint(height)
		preferredWidth = uint(float32(originalWidth * height) / float32(originalHeight))
	}

	resizedImage := resize.Resize(preferredWidth, preferredHeight, original, resize.NearestNeighbor)
	return &resizedImage
}