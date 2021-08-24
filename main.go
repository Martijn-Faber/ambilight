package main

import (
	"time"

	"github.com/Martijn-Faber/ambilight/config"
	"github.com/Martijn-Faber/ambilight/internal/block"
	"github.com/Martijn-Faber/ambilight/internal/protocol/wled"
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

	for {
		img, err := screenshot.CaptureScreen()

		if err != nil {
			panic(err)
		}

		clr := block.GetBlocks(img, blkSizeX, blkSizeY, numBlkX, numBlkY, scrWidth)

		err = wled.Send(clr)

		if err != nil {
			panic(err)
		}

		time.Sleep(time.Duration(config.RefTimeMs) * time.Millisecond)
	}
}
