package main

import (
	"io/ioutil"
	"testing"
)

func BenchmarkCompressionGzip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GzipCompression(ioutil.Discard)
	}
}

func BenchmarkCompressionBrotli4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BrotliCompression(ioutil.Discard, 4)
	}
}

func BenchmarkCompressionBrotli11(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BrotliCompression(ioutil.Discard, 11)
	}
}

func BenchmarkDecompressionGzip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GzipDecompression()
	}
}

func BenchmarkDecompressionBrotli4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Brotli4Decompression()
	}
}

func BenchmarkDecompressionBrotli11(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Brotli11Decompression()
	}
}
