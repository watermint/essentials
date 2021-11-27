package edesktop

import (
	"github.com/watermint/essentials/eidiom"
)

type Desktop interface {
	// Open Launches the associated application to open the file.
	// Note: the parameter type will be replaced by efs.Path in the future.
	Open(path string) OpenOutcome
}

type OpenOutcome interface {
	eidiom.Outcome

	IsOpenFailure() bool
	IsOperationUnsupported() bool
}
