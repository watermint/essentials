package edesktop

import (
	"path/filepath"
)

func CurrentDesktop() Desktop {
	return &desktopImpl{}
}

type desktopImpl struct {
}

func (z desktopImpl) Open(path string) OpenOutcome {
	// desktopOpen expects local path in file system format
	return desktopOpen(filepath.Clean(path))
}
