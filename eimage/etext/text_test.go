package etext

import (
	"github.com/golang/freetype/truetype"
	"github.com/watermint/essentials/eimage/ecolor"
	"github.com/watermint/essentials/eimage/egeom"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomedium"
	"image"
	"image/draw"
	"image/png"
	"os"
	"testing"
)

func TestDraw(t *testing.T) {
	text := "Hello your go World"
	fill, _ := ecolor.ParseColor("marker(b18)")
	textFill, _ := ecolor.ParseColor("marker(w00)")
	padding := egeom.NewPaddingRatio(0.5, 0.5)
	position, _ := egeom.ParsePosition("center")

	fontSize := 36
	width := 640
	height := 400

	r := image.Rect(0, 0, width, height)
	img := image.NewRGBA(r)
	draw.Draw(img, img.Bounds(), image.NewUniform(fill), image.Point{}, draw.Src)

	textFont, _ := truetype.Parse(gomedium.TTF)
	face := truetype.NewFace(textFont, &truetype.Options{
		Size:              float64(fontSize),
		DPI:               0,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	})

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(textFill),
		Face: face,
	}
	boundString, _ := d.BoundString(text)

	strPosition := position.Locate(egeom.NewRectangleImage(img.Bounds()), egeom.NewRectangleFixed26(boundString), padding)
	d.Dot = strPosition.Fixed26()
	d.DrawString(text)

	f, err := os.Create("/tmp/test_draw.png")
	if err != nil {
		t.Error(err)
		return
	}
	if err := png.Encode(f, img); err != nil {
		t.Error(err)
	}
	_ = f.Close()
}
