# compression-comparison

A small tool to compare compression with [Brotli](https://github.com/google/brotli) and
[gzip](https://en.m.wikipedia.org/wiki/Gzip).

The idea is to see how gzip and Brotli compare on your own data. To use it, replace `input` with an
example of the data you'd want to compress, e.g. a JSON file with a response from your API. Then run

```
go run .
```

to see compression ratios, and

```
go test -bench .
```

to compare compression and decompression speed.

Brotli has a quality setting that lets you trade off speed and effectiveness. The benchmarks include
quality 4, which should work well for on-the-fly encryption, and quality 11, which will give the
best compression ratio.


## Prerequisites

You'll need

* [Go](https://golang.org/dl/), of course.
* The Brotli library, including header files. On Ubuntu, `sudo apt install libbrotli-dev` should do
  the trick.
