package main

import (
	"time"

	"github.com/Martijn-Faber/ambilight/config"
	"github.com/Martijn-Faber/ambilight/internal/block"
	"github.com/Martijn-Faber/ambilight/internal/protocol/wled"
	"github.com/Martijn-Faber/ambilight/pkg/color"
	"github.com/Martijn-Faber/ambilight/pkg/screenshot"
	"github.com/Martijn-Faber/ambilight/util"
)

func main() {
	config, err := config.Load()

	if err != nil {
		panic(err)
	}

	wled.Init(config.WledIp, config.WledPort)

	scrWidth, scrHeight, _ := util.GetScreenInfo()
	blkSizeX, blkSizeY := block.CalculateBlockSize(scrWidth, scrHeight, config.NumLedsX, config.NumLedsY)
	numBlkX, numBlkY := block.CalculateNumBlocks(scrWidth, scrHeight, blkSizeX, blkSizeY)

	var prevClrs []color.Rgb

	for {
		img, err := screenshot.CaptureScreen()

		if err != nil {
			panic(err)
		}

		nextClrs := block.GetBlocks(img, blkSizeX, blkSizeY, numBlkX, numBlkY, scrWidth)

		if len(prevClrs) == 0 {
			prevClrs = nextClrs
		}

		for i := 0; i < config.InterpolSteps; i++ {
			clrs := color.InterpolateColors(prevClrs, nextClrs, (i + 1), config.InterpolSteps)
			err = wled.Send(clrs)

			if err != nil {
				panic(err)
			}

			prevClrs = nextClrs
			time.Sleep(time.Duration(config.RefTimeMs/config.InterpolSteps) * time.Millisecond)
		}
	}
}
