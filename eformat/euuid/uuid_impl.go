package euuid

import (
	"crypto/rand"
	"essentials/eformat/ehex"
)

func NewV4() UUID {
	u := make([]byte, 16)

	// Generate random sequence
	_, err := rand.Read(u)
	if err != nil {
		panic(err)
	}

	// Copy from Google's impl.
	// ---
	// Copyright 2016 Google Inc.  All rights reserved.
	// Use of this source code is governed by a BSD-style
	// license that can be found in the LICENSE file.
	u[6] = (u[6] & 0x0f) | 0x40 // Version 4
	u[8] = (u[8] & 0x3f) | 0x80 // Variant is 10

	return &uuidV4{
		u: u,
	}
}

type uuidV4 struct {
	u []byte
}

func (z uuidV4) String() string {
	return ehex.ToHexString(z.u[0:4]) + "-" +
		ehex.ToHexString(z.u[4:6]) + "-" +
		ehex.ToHexString(z.u[6:8]) + "-" +
		ehex.ToHexString(z.u[8:10]) + "-" +
		ehex.ToHexString(z.u[10:16])
}

func (z uuidV4) Urn() string {
	return "urn:uuid:" + z.String()
}
