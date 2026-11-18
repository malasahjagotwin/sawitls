# sawitls

**TLS for Go that can present a browser-like ClientHello.**

[![Go](https://img.shields.io/badge/go-1.24%2B-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev)
[![License](https://img.shields.io/badge/license-BSD--3--Clause-blue?style=flat-square)](LICENSE)
[![Release](https://img.shields.io/badge/release-v0.1.0-2ea44f?style=flat-square)](https://github.com/malasahjagotwin/sawitls/tags)
[![Reference](https://img.shields.io/badge/reference-godoc-007d9c?style=flat-square)](https://pkg.go.dev/github.com/malasahjagotwin/sawitls)

Standard `crypto/tls` sends one fixed ClientHello, which is easy to recognise.
`sawitls` keeps the same API but lets you choose the handshake fingerprint
(Chrome, Firefox, Safari, randomised, or your own), and exposes the handshake
in both directions.

## Install

```sh
go get github.com/malasahjagotwin/sawitls@v0.1.0
```

## Quickstart

```go
package main

import (
	"fmt"
	"net"

	tls "github.com/malasahjagotwin/sawitls"
)

func main() {
	raw, err := net.Dial("tcp", "example.com:443")
	if err != nil {
		panic(err)
	}
	uconn := tls.UClient(raw, &tls.Config{ServerName: "example.com"}, tls.HelloChrome_Auto)
	if err := uconn.Handshake(); err != nil {
		panic(err)
	}
	defer uconn.Close()

	fmt.Println("fingerprint:", uconn.ClientHelloID.Str())
}
```

Pick `tls.HelloChrome_Auto` for the current Chrome profile, `tls.HelloRandomized`
for a shuffled-but-valid hello, or `tls.HelloCustom` when you build the
ClientHello yourself with `BuildHandshakeState`.

## Features

- **ClientHello control**: cipher suites, extensions, curves, signature algorithms.
- **Ready-made fingerprints**: Chrome, Firefox, Safari, iOS, Android, randomised.
- **Low-level access**: read/write every part of the hello, inspect server hello.
- **Same shape as `crypto/tls`**: `Dial`, `UClient`, `UConn`, `Config`.
- **Session tickets**: fake or real, resumption still works.

## Layout

| Path | Purpose |
| --- | --- |
| `sawitls` | client, handshake, config |
| `sawitls/dicttls` | certificate dictionaries |
| `sawitls/internal` | crypto and helper packages |
| `sawitls/fipsonly` | FIPS-only build |
| `sawitls/examples` | runnable examples |

## Requirements

Go 1.24 or newer.

## License

BSD-3-Clause. See [LICENSE](LICENSE).
