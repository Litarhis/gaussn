package pic // import "github.com/litarhis/gaussn/internalPackages/pic"

import (
	"bytes"
	"image"
	"image/png"
)

func Show(f func(int, int) [][]uint8, width int, height int) []byte {
	data := f(width, height)
	m := image.NewNRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = 0   //----------
			m.Pix[i+1] = 0 // Greyscale
			m.Pix[i+2] = 0 //----------

			m.Pix[i+3] = v // density
		}
	}
	return Encode2Png(m)
}

func Encode2Png(m image.Image) []byte {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}
