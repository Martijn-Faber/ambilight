package block

import (
	"image"

	"github.com/Martijn-Faber/ambilight/pkg/color"
)

func GetBlocks(img *image.RGBA, blkSizeX int, blkSizeY int, numBlkX int, numBlkY int, scrWidth int) []color.Rgb {
	var clr []color.Rgb

	// bottom right to top right
	for blkY := (numBlkY - 1); blkY >= 0; blkY-- {
		clr = append(clr, GetAverageBlockColor(img, (numBlkX-1), blkY, blkSizeX, blkSizeY, scrWidth))
	}

	// top right to left top
	for blkX := (numBlkX - 1); blkX >= 0; blkX-- {
		clr = append(clr, GetAverageBlockColor(img, blkX, 0, blkSizeX, blkSizeY, scrWidth))
	}

	// left top to bottom left
	for blkY := 0; blkY < numBlkY; blkY++ {
		clr = append(clr, GetAverageBlockColor(img, 0, blkY, blkSizeX, blkSizeY, scrWidth))
	}

	return clr
}
