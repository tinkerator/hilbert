// Program curve renders Hilbert curves in the form of a PNG file.
package main

import (
	"flag"
	"image"
	"image/color"
	"log"

	"github.com/llgcode/draw2d/draw2dimg"
	"zappem.net/pub/math/hilbert"
)

var (
	wr    = flag.String("write", "", "write png to this named file")
	n     = flag.Int("n", 4, "scale power of 2 for curve nodes")
	width = flag.Int("width", 256, "png image width")
)

func main() {
	flag.Parse()

	if *wr == "" {
		log.Fatal("usage: curve --write=<file.png>")
	}

	img := image.NewRGBA(image.Rect(0, 0, *width, *width))
	gc := draw2dimg.NewGraphicContext(img)
	gc.ClearRect(0, 0, *width, *width)
	gc.SetStrokeColor(color.RGBA{0, 0, 255, 255})

	w := uint(1) << uint(*n+1)
	var d = float64(uint(*width)/w) * .5
	var X, Y float64 = d, d

	for s := uint(1); ; s++ {
		x, y, err := hilbert.StepXY(uint(*n), s)
		if err != nil {
			break
		}
		gc.BeginPath()
		gc.MoveTo(X, Y)
		X, Y = d+float64(x*uint(*width)/w), d+float64(y*uint(*width)/w)
		gc.LineTo(X, Y)
		gc.Stroke()
	}

	if err := draw2dimg.SaveToPngFile(*wr, img); err != nil {
		log.Fatalf("unable to save %q: %v", *wr, err)
	}
}
