package block

import (
	"image"
)

func GetAverageBlockColor(img *image.RGBA, blkX int, blkY int, blkSizeX int, blkSizeY int, scrWidth int) Rgb {
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
			var pixelIndex = lineStart + x * bytesPerPixel

			rSum += float32(img.Pix[pixelIndex + 0 ])
			gSum += float32(img.Pix[pixelIndex + 1 ])
			bSum += float32(img.Pix[pixelIndex + 2 ])
			totalPix++
		}
	}

	var rAvg = rSum / totalPix
	var gAvg = gSum / totalPix
	var bAvg = bSum / totalPix

	for y := blkStartY; y < (blkStartY + blkSizeY); y++ {
		var lineStart = y * bytesPerLine
		for x := blkStartX; x < (blkStartX + blkSizeX); x++ {
			var pixelIndex = lineStart + x * bytesPerPixel

			img.Pix[pixelIndex + 0 ] = uint8(rAvg)
			img.Pix[pixelIndex + 1 ] = uint8(gAvg)
			img.Pix[pixelIndex + 2 ] = uint8(bAvg)
			img.Pix[pixelIndex + 3 ] = 255	
		}
	}

	return Rgb{
		R: byte(rAvg),
		G: byte(gAvg),
		B: byte(bAvg),
	}
}
