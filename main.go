package main

import (
	"fmt"
	"time"

	"github.com/Martijn-Faber/Ambilight/config"
	"github.com/Martijn-Faber/Ambilight/internal/block"
	"github.com/Martijn-Faber/Ambilight/internal/protocol/wled"
	"github.com/Martijn-Faber/Ambilight/pkg/screenshot"
	"github.com/Martijn-Faber/Ambilight/util"
)

func main() {
	config, err := config.Load()

	if err != nil {
		panic(err)
	}

	wled.Init(config.WledIp, config.WledPort)
	scrWidth, scrHeight, _ := util.GetScreenInfo()

	for {
		img, err := screenshot.CaptureScreen()

		if err != nil {
			panic(err)
		}
	
		blkSizeX, blkSizeY := block.CalculateBlockSize(scrWidth, scrHeight, config.NumLedsX, config.NumLedsY)
		numBlkX, numBlkY := block.CalculateNumBlocks(scrWidth, scrHeight, blkSizeX, blkSizeY)
	
		clr := block.GetBlocks(img, blkSizeX, blkSizeY, numBlkX,  numBlkY, scrWidth)

		err = wled.Send(clr)

		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(time.Duration(config.RefTimeMs) * time.Millisecond)
	}
}