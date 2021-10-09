package etext

import "testing"

func TestNewTrueType(t *testing.T) {
	for _, df := range DefaultFonts {
		b, a := df.BoundString("Hello")
		if b.Width() < 1 || b.Height() < 1 || a < 1 {
			t.Error(b, a)
		}
	}
}
