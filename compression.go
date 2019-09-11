package main

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"

	"github.com/google/brotli/go/cbrotli"
)

var input, gzipped, brotli4ed, brotli11ed, output []byte

func init() {
	var err error
	input, err = ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	GzipCompression(buf)
	gzipped = buf.Bytes()

	buf = new(bytes.Buffer)
	BrotliCompression(buf, 4)
	brotli4ed = buf.Bytes()

	buf = new(bytes.Buffer)
	BrotliCompression(buf, 11)
	brotli11ed = buf.Bytes()

	output = make([]byte, len(input))
}

func GzipCompression(w io.Writer) {
	writer := gzip.NewWriter(w)
	defer writer.Close()
	_, err := writer.Write(input)
	if err != nil {
		panic(err)
	}
}

func BrotliCompression(w io.Writer, quality int) {
	options := cbrotli.WriterOptions{Quality: quality}
	writer := cbrotli.NewWriter(w, options)
	defer writer.Close()
	_, err := writer.Write(input)
	if err != nil {
		panic(err)
	}
}

func GzipDecompression() {
	reader, err := gzip.NewReader(bytes.NewReader(gzipped))
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	_, err = reader.Read(output)
	if err != nil {
		panic(err)
	}
}

func Brotli4Decompression() {
	reader := cbrotli.NewReader(bytes.NewReader(brotli4ed))
	defer reader.Close()
	_, err := reader.Read(output)
	if err != nil {
		panic(err)
	}
}

func Brotli11Decompression() {
	reader := cbrotli.NewReader(bytes.NewReader(brotli11ed))
	defer reader.Close()
	_, err := reader.Read(output)
	if err != nil {
		panic(err)
	}
}
