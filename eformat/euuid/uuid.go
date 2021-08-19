package euuid

import (
	"fmt"
	"regexp"
)

type UUID interface {
	fmt.Stringer

	// Urn returns URN form like `urn:uuid:123e4567-e89b-12d3-a456-426655440000`.
	Urn() string
}

const (
	uuidRePattern = `^[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$`
)

var (
	uuidRe = regexp.MustCompile(uuidRePattern)
)

func IsUUID(uuid string) bool {
	return uuidRe.MatchString(uuid)
}
