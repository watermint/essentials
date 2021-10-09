package etext

import (
	"github.com/golang/freetype/truetype"
	"github.com/watermint/essentials/egraphic/egeom"
	"github.com/watermint/essentials/eidiom/eoutcome"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/gobolditalic"
	"golang.org/x/image/font/gofont/goitalic"
	"golang.org/x/image/font/gofont/gomedium"
	"golang.org/x/image/font/gofont/gomediumitalic"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/gofont/gomonobold"
	"golang.org/x/image/font/gofont/gomonobolditalic"
	"golang.org/x/image/font/gofont/gomonoitalic"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/gofont/gosmallcaps"
	"golang.org/x/image/font/gofont/gosmallcapsitalic"
)

const (
	DefaultFontSize = 12
)

var (
	GoFontRegular         = MustNewTrueTypeParse(goregular.TTF)
	GoFontItalic          = MustNewTrueTypeParse(goitalic.TTF)
	GoFontBold            = MustNewTrueTypeParse(gobold.TTF)
	GoFontBoldItalic      = MustNewTrueTypeParse(gobolditalic.TTF)
	GoFontMedium          = MustNewTrueTypeParse(gomedium.TTF)
	GoFontMediumItalic    = MustNewTrueTypeParse(gomediumitalic.TTF)
	GoFontMono            = MustNewTrueTypeParse(gomono.TTF)
	GoFontMonoItalic      = MustNewTrueTypeParse(gomonoitalic.TTF)
	GoFontMonoBold        = MustNewTrueTypeParse(gomonobold.TTF)
	GoFontMonoBoldItalic  = MustNewTrueTypeParse(gomonobolditalic.TTF)
	GoFontSmallCaps       = MustNewTrueTypeParse(gosmallcaps.TTF)
	GoFontSmallCapsItalic = MustNewTrueTypeParse(gosmallcapsitalic.TTF)

	DefaultFonts = []Font{
		GoFontRegular,
		GoFontItalic,
		GoFontBold,
		GoFontBoldItalic,
		GoFontMedium,
		GoFontMediumItalic,
		GoFontMono,
		GoFontMonoItalic,
		GoFontMonoBold,
		GoFontMonoBoldItalic,
		GoFontSmallCaps,
		GoFontSmallCapsItalic,
	}
)

type Font interface {
	WithSize(size int) Font

	BoundString(text string) (bound egeom.Rectangle, advance int)

	// Size returns vertical font size in pixel.
	Size() int

	Face() font.Face
}

func MustNewTrueTypeParse(fontData []byte) Font {
	f, oc := NewTrueTypeParse(fontData)
	if oc.IsError() {
		panic(oc.String())
	}
	return f
}

func NewTrueTypeParse(fontData []byte) (f Font, oc eoutcome.ParseOutcome) {
	ttf, err := truetype.Parse(fontData)
	if err != nil {
		return nil, eoutcome.NewParseInvalidFormat(err.Error())
	}
	return NewTrueType(ttf), eoutcome.NewParseSuccess()
}

func NewTrueType(f *truetype.Font) Font {
	return &ttfImpl{
		ttf:  f,
		size: DefaultFontSize,
	}
}

type ttfImpl struct {
	ttf  *truetype.Font
	size int
}

func (z ttfImpl) Size() int {
	return z.size
}

func (z ttfImpl) WithSize(size int) Font {
	z.size = size
	return z
}

func (z ttfImpl) BoundString(text string) (bound egeom.Rectangle, advance int) {
	d := &font.Drawer{
		Face: z.Face(),
	}
	b26, a26 := d.BoundString(text)
	return egeom.NewRectangleFixed26(b26), a26.Round()
}

func (z ttfImpl) Face() font.Face {
	return truetype.NewFace(
		z.ttf,
		&truetype.Options{
			Size: float64(z.size),
		},
	)
}
