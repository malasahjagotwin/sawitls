//[[:space:]]*Copyright (c) 2026 malasahjagotwin. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build boringcrypto

package fipsonly

import (
	"crypto/tls/internal/fips140tls"
	"testing"
)

func Test(t *testing.T) {
	if !fips140tls.Required() {
		t.Fatal("fips140tls.Required() = false, must be true")
	}
}
