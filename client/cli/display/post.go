package display

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"io"
	"math"
	"strings"

	"github.com/tomnomnom/xtermcolor"
)

type ImgDisplay struct{}

func (ImgDisplay) GetDisplayString(termWidth int, s io.Reader) (strings.Builder, error) {
	im, _, err := image.Decode(s)
	if err != nil {
		return strings.Builder{}, err
	}

	xM := 1
	yM := 2 //terminal cell(pixel) is approximately 2 times taller than wider
	if im.Bounds().Max.X > termWidth {
		multiplier := int(math.Ceil(float64(im.Bounds().Max.X) / float64(termWidth)))
		xM = multiplier
		yM *= multiplier
	}

	var bd strings.Builder
	size := float32(xM * yM)
	for y := 0; y+yM < im.Bounds().Max.Y; y += yM {
		for x := 0; x+xM < im.Bounds().Max.X; x += xM {
			rgba := avgColor(y, yM, x, xM, im, size)

			// all the magic in the next 2 rows
			n := xtermcolor.FromColor(rgba)
			bd.Write([]byte(fmt.Sprintf("\033[48;5;%dm \033[0m", n)))
		}
		bd.Write([]byte("\n"))
	}

	return bd, nil
}

func avgColor(y int, yM int, x int, xM int, im image.Image, size float32) color.RGBA64 {
	var r, g, b, a float32
	for i := y; i < y+yM; i++ {
		for j := x; j < x+xM; j++ {
			p := im.At(j, i)
			pr, pg, pb, pa := p.RGBA()
			r += float32(pr) / size
			g += float32(pg) / size
			b += float32(pb) / size
			a += float32(pa) / size
		}
	}

	return color.RGBA64{
		R: uint16(r),
		G: uint16(g),
		B: uint16(b),
		A: uint16(a),
	}
}
