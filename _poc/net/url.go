package net

import "fmt"

type Url interface {
	fmt.Stringer

	// Authority component
	Authority() string

	// Fragment component
	Fragment() string

	// Host component
	Host() string

	// Path component
	Path() string

	// Port number
	Port() int

	// Query component
	Query() string

	// Scheme component
	Scheme() string

	// UserInfo component
	UserInfo() UserInfo
}

type UserInfo interface {
}
