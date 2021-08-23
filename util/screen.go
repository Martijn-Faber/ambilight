package util

import (
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

func GetScreenInfo() (scrWidth int, scrHeight int, root xproto.Window) {
	c, err := xgb.NewConn()

	// TODO: return error instead of handeling it here
	if err != nil {
		panic(err)
	}

	defer c.Close()

	screen := xproto.Setup(c).DefaultScreen(c)
	scrWidth = int(screen.WidthInPixels)
	scrHeight = int(screen.HeightInPixels)
	root = screen.Root
	
	return
}