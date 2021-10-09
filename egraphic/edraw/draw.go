package edraw

import (
	"github.com/watermint/essentials/egraphic/ecolor"
	"github.com/watermint/essentials/egraphic/egeom"
	"github.com/watermint/essentials/egraphic/eimage"
	"github.com/watermint/essentials/egraphic/etext"
	"golang.org/x/image/font"
	"image"
	"image/draw"
)

type Draw interface {
	// FillRectangle a rectangle with the color
	FillRectangle(rect egeom.Rectangle, color ecolor.Color)

	DrawString(pos egeom.Point, text string, style etext.Style)
}

func NewImageDrawer(img eimage.Image) Draw {
	return &drawImpl{
		img: img,
	}
}

type drawImpl struct {
	img eimage.Image
}

func (z drawImpl) FillRectangle(rect egeom.Rectangle, color ecolor.Color) {
	draw.Draw(z.img.GoImageRGBA(), rect.ImageRect(), image.NewUniform(color), rect.TopLeft().ImagePoint(), draw.Src)
}

func (z drawImpl) DrawString(pos egeom.Point, text string, style etext.Style) {
	dr := &font.Drawer{
		Dst:  z.img.GoImageRGBA(),
		Src:  image.NewUniform(style.Color()),
		Face: style.Font().Face(),
	}
	style.Layout(text, pos, func(s string, p egeom.Point) {
		dr.Dot = p.Add(0, dr.Face.Metrics().Height.Round()).Fixed26()
		dr.DrawString(s)
	})
}
