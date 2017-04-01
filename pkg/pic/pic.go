package pic // import "gaussn/pkg/pic"

import (
	"bytes"
	"image"
	"image/png"
)

func Show(f func(int, int) [][]uint8) []byte {
	const (
		dx = 256
		dy = 256
	)
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	return ShowImage(m)
}

func ShowImage(m image.Image) []byte {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}
