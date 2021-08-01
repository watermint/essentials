package desktop

import (
	"essentials/_poc/fs"
	"essentials/_poc/net"
)

type Desktop interface {
	// Browse Launches the default browser to display a URL. Wait for complete.
	Browse(url net.Url) BrowseResult

	// Open Launches the associated application to open the file. Wait for complete.
	Open(file fs.Path) OpenResult
}

type DesktopError interface {
	error
}

type BrowseResult interface {
	IfSuccess(f func()) BrowseResult
	IfFailure(f func(err DesktopError)) BrowseResult
}

type OpenResult interface {
	IfSuccess(f func()) OpenResult
	IfFailure(f func(err DesktopError)) OpenResult
}
