package block

import (
	"image"

	"github.com/Martijn-Faber/ambilight/pkg/color"
)

func GetAverageBlockColor(img *image.RGBA, blkX int, blkY int, blkSizeX int, blkSizeY int, scrWidth int) color.Rgb {
	var rSum float32
	var gSum float32
	var bSum float32
	var totalPix float32

	bytesPerPixel := 4
	bytesPerLine := bytesPerPixel * scrWidth

	blkStartX := blkX * blkSizeX
	blkStartY := blkY * blkSizeY

	for y := blkStartY; y < (blkStartY + blkSizeY); y++ {
		lineStart := y * bytesPerLine

		for x := blkStartX; x < (blkStartX + blkSizeX); x++ {
			var pixelIndex = lineStart + x*bytesPerPixel

			bSum += float32(img.Pix[pixelIndex+0])
			gSum += float32(img.Pix[pixelIndex+1])
			rSum += float32(img.Pix[pixelIndex+2])
			totalPix++
		}
	}

	var rAvg = rSum / totalPix
	var gAvg = gSum / totalPix
	var bAvg = bSum / totalPix

	return color.Rgb{R: rAvg, G: gAvg, B: bAvg}
}
