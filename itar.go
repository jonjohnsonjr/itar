package itar

import (
	"archive/tar"
	"errors"
	"io"
	"iter"
)

func Etor(tr *tar.Reader) iter.Seq2[*tar.Header, error] {
	return func(yield func(*tar.Header, error) bool) {
		for {
			hdr, err := tr.Next()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return
				}
			}

			if !yield(hdr, err) {
				return
			}
		}
	}
}
