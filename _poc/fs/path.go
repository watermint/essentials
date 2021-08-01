package fs

import (
	"fmt"
)

type Path interface {
	fmt.Stringer

	// Basename Returns the last element of the path.
	Basename() string

	// Parent path.
	Parent() ParentResult

	// Child path.
	Child(name ...string) (p Path)

	// Exists True when the path exists on the file system.
	Exists() ExistsResult

	// Relative Calc relative path
	Relative(other Path) RelativeResult
}

type ParentResult interface {
	IfSuccess(f func(parent Path)) ParentResult
	IfFailure(f func(err ParentError)) ParentResult
}

type RelativeResult interface {
	IfSuccess(f func(rel Path)) RelativeResult
	IfFailure(f func(err PathError)) RelativeResult
}

type ParentError interface {
	PathError

	IsNoParent() bool
	IsNotVisible() bool
	IsNotAccessible() bool
}

type PathError interface {
	error

	IsInvalidPath() bool
}
