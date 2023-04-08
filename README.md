# hilbert

## Overview

This package provides a one method package to help render Hilbert curves.

Automated documentation for this Go package is available from
[![Go Reference](https://pkg.go.dev/badge/zappem.net/pub/math/hilbert.svg)](https://pkg.go.dev/zappem.net/pub/math/hilbert).

While the `hilbert` package has no dependencies beyond the standard Go
packages, we include an example binary, `example/curve.go`, that
serves as a demonstration of using the package and includes the
generation of a png file with a rendered Hilbert curve. This example
uses the [`draw2d`](https://pkg.go.dev/github.com/llgcode/draw2d)
package to render the curve.

To try this example:
```
$ go mod tidy
$ go build example/curve.go
$ ./curve --write=example.png
```

The generated PNG looks like this:

![sample output of the curve program](sample.png)

Supplying different `--n` values to the `curve` program will generate
different scales of the space filling curve.

## License info

The `hilbert` package is distributed with the same BSD 3-clause license
as that used by [golang](https://golang.org/LICENSE) itself.

## Reporting bugs and feature requests

Use the [github `hilbert` bug
tracker](https://github.com/tinkerator/hilbert/issues).
