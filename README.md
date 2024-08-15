# `itar`

This provides an iterator for `archive/tar`.

```go
package main

import (
	"fmt"
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
```