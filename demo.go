package main

import (
	"encoding/hex"
	"fmt"
	"internal/hashwriter"
	"os"
)

func main() {

	h := hashwriter.NewHashWriter(os.Stdout)
	h.Write([]byte("hello\n\n"))

	b := h.Sum([]byte{})
	fmt.Printf("SHA256 checksum %s\n\n", hex.EncodeToString(b))

}
