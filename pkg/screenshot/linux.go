package screenshot

import (
	"image"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/Martijn-Faber/Ambilight/util"
)

var conn *xgb.Conn
var scrWidth int
var scrHeight int
var root xproto.Window

func init() {
	c, err := xgb.NewConn()

	if err != nil {
		panic(err)
	}

	conn = c

	scrW, scrH, r := util.GetScreenInfo()

	scrWidth = scrW
	scrHeight = scrH
	root = r
}

func CaptureScreen() (*image.RGBA, error) {
	r := image.Rect(0, 0, int(scrWidth), int(scrHeight))
	i, err := xproto.GetImage(conn, xproto.ImageFormatZPixmap, xproto.Drawable(root), int16(r.Min.X), int16(r.Min.Y), uint16(scrWidth), uint16(scrHeight), 0xffffffff).Reply()

	if err != nil {
		return nil, err
	}

	data := i.Data
	for i := 0; i < len(data); i += 4 {
		data[i], data[i+2], data[i+3] = data[i+2], data[i], 255
	}

	return &image.RGBA{data, 4 * r.Dx(), r}, err
}