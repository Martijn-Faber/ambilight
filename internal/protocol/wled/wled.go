package wled

import (
	"net"

	"github.com/Martijn-Faber/ambilight/pkg/color"
)

var conn net.Conn

func Init(ip string, port int) {
	addr := net.UDPAddr{
		IP:   net.ParseIP(ip),
		Port: port,
	}

	c, err := net.Dial("udp", addr.String())

	if err != nil {
		panic(err)
	}

	conn = c
}

func Send(clrs []color.Rgb) error {
	var buf []byte

	// protocol type
	// time to wait in seconds
	buf = append(buf, 1, 10)

	for i, clr := range clrs {
		// LED index
		// red value
		// green value
		// blue value
		buf = append(buf, byte(i), clr.R, clr.G, clr.B)
	}

	_, err := conn.Write(buf)

	if err != nil {
		return err
	}

	return err
}
