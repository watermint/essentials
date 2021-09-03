package eidiom

import "fmt"

// Outcome is the alternative for the `error`. Outcome instance must not be nil.
// And Outcome implementation should implement specific error cases with the prefix `IsXxx`.
// For example: file operation.
//
//   type FileCreateOutcome interface {
//     Outcome
//     IsPermissionDenied() bool
//     IsInvalidPath() bool
//     IsOperationNotAllowed() bool
//   }
//
// Consumer can handle errors like below.
//
//    f, out := eio.Create("/path/to/create")
//    switch {
//    case out.IsOk():
//        // success
//    case out.IsPermissionDenied(), out.IsOperationNotAllowed():
//        // handle permission issue
//    case out.IsInvalidPath():
//        // handle path issue
//    default:
//        // handle other errors
//    }
type Outcome interface {
	// Stringer Outcome instance returns empty string if an operation succeed, otherwise returns an error string.
	fmt.Stringer

	// IsOk Returns true if an operation succeed.
	IsOk() bool

	// IfOk Perform f if an operation succeed, otherwise does nothing.
	IfOk(f func())

	// IsError Returns true if an operation got an error.
	IsError() bool

	// IfError Perform f if an operation got an error, otherwise does nothing.
	IfError(f func() Outcome) Outcome

	// Cause Return Outcome as an error instance if an operation got an error, otherwise returns nil.
	Cause() error
}
