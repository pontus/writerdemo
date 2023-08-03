package hashwriter

import (
	"crypto/sha256"
	"errors"
	"hash"
	"io"
)

type HashWriter struct {
	w io.Writer
	h hash.Hash
}

func NewHashWriter(w io.Writer) *HashWriter {
	r := &HashWriter{w: w}
	r.h = sha256.New()
	return r
}

func (hw *HashWriter) Sum(b []byte) []byte {
	return hw.h.Sum(b)
}

func (hw *HashWriter) Reset() {
	hw.h.Reset()
}

func (hw *HashWriter) Write(data []byte) (int, error) {
	n, err := hw.h.Write(data)

	if err != nil {
		return n, err
	}

	if n != len(data) {
		return n, errors.New("Wrote less than expected to hash")
	}

	written := 0

	for written < len(data) {

		nx, err := hw.w.Write(data)
		written += nx
		if err != nil {
			return written, err
		}
	}

	return written, nil
}
