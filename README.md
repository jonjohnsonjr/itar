# `itar`

This provides an iterator for `archive/tar`.

```go
package main

import (
	"archive/tar"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/jonjohnsonjr/itar"
)

func main() {
	tr := tar.NewReader(os.Stdin)
	for hdr, err := range itar.Etor(tr) {
		if err != nil {
			panic(err)
		}

		fmt.Println(hdr.FileInfo().(fmt.Stringer).String())

		if n, err := io.Copy(io.Discard, tr); err != nil {
			panic(err)
		} else {
			fmt.Printf("read %d bytes\n", n)
		}
	}
}

func before() {
	tr := tar.NewReader(os.Stdin)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			panic(err)
		}

		fmt.Println(hdr.FileInfo().(fmt.Stringer).String())

		if n, err := io.Copy(io.Discard, tr); err != nil {
			panic(err)
		} else {
			fmt.Printf("read %d bytes\n", n)
		}
	}
}
```
