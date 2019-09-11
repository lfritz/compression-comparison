package main

import (
	"bytes"
	"fmt"
)

func main() {
	inputLen := len(input)
	fmt.Printf("%-20s %6d bytes\n", "Input:", inputLen)

	buf := new(bytes.Buffer)
	GzipCompression(buf)
	gzipLen := buf.Len()
	fmt.Printf("%-20s %6d bytes\n", "Gzip:", gzipLen)

	buf = new(bytes.Buffer)
	BrotliCompression(buf, 4)
	brotli4Len := buf.Len()
	fmt.Printf("%-20s %6d bytes\n", "Brotli (quality 4):", brotli4Len)

	buf = new(bytes.Buffer)
	BrotliCompression(buf, 11)
	brotli11Len := buf.Len()
	fmt.Printf("%-20s %6d bytes\n", "Brotli (quality 11):", brotli11Len)
}
