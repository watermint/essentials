//go:build windows
// +build windows

package ewindows

import "syscall"

type Kernel interface {
	Call(procName string, args ...uintptr) (r1, r2 uintptr, lastErr, resolveErr error)
}

var (
	Kernel32 Kernel = &kernelWrapper{}
)

type kernelWrapper struct {
}

func (z kernelWrapper) Call(procName string, args ...uintptr) (r1, r2 uintptr, lastErr, resolveErr error) {
	k32, resolveErr := syscall.LoadDLL("kernel32")
	if resolveErr != nil {
		return 0, 0, nil, resolveErr
	}
	proc, resolveErr := k32.FindProc(procName)
	if resolveErr != nil {
		return 0, 0, nil, resolveErr
	}

	r1, r2, lastErr = proc.Call(args...)
	return r1, r2, lastErr, nil
}
