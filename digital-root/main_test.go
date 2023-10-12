package main

import (
	"testing"
)

func TestDigitalRootWithString(t *testing.T) {
	input := 1234
	want := 1
	got := DigitalRootWithString(input)
	if want != got {
		t.Errorf("got DigitalRootWithString(%d) = %d, wanted: %d", input, got, want)
	}
}

func TestDigitalRootWithModulo(t *testing.T) {
	input := 1234
	want := 1
	got := DigitalRootWithModulo(input)
	if want != got {
		t.Errorf("got DigitalRootWithModulo(%d) = %d, wanted: %d", input, got, want)
	}
}

func BenchmarkDigitalRootWithString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DigitalRootWithString(1234)
	}
}

func BenchmarkDigitalRootWithModulo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DigitalRootWithModulo(1234)
	}
}

// -run flag for regex on test functions
// -bench flag for regex on benchmark functions
// -benchtime=10s specifies time of execution
// -benchtime=100x specifies b.N (100 iterations)

// to run only all benchmark functions - use below
// go test -run=^$ -bench=. -benchmem

// to run only the tests functions - use below
// go test -v

// to run modulo test function only - use below
// go test -run=^(DigitalRootWithModulo)$ -v
// go test -run=^TestDigitalRootWithModulo$ -v

// to run modulo test and all benchmarks - use below
// go test -run=^TestDigitalRootWithModulo$ -bench=. -benchmem -v

// to run only string test and its benchmark functions - use below
// go test -run ^TestDigitalRootWithString$ -bench ^BenchmarkDigitalRootWithModulo$ -benchmem -v

// to run all tests (run regex or add -v) with specific benchmark
// go test -bench ^BenchmarkDigitalRootWithModulo$ -benchmem -v
// go test -run . -bench ^BenchmarkDigitalRootWithModulo$ -benchmem -v

// to run one specific benchmark
// go test -bench ^BenchmarkDigitalRootWithModulo$ -benchmem
// go test -run=Benchmark -bench=DigitalRootWithModulo$ -benchmem

/*/ ------------------------- output -------------------------
// benchmarks results only.

[13:13:59] {nxos-geek}:~$ go test -run=^$ -bench=. -benchmem -benchtime=20000000x
goos: windows
goarch: amd64
pkg: github.com/jeamon/digital-root-algorithms
BenchmarkDigitalRootWithString-4        20000000               108 ns/op               4 B/op          1 allocs/op
BenchmarkDigitalRootWithModulo-4        20000000                13.5 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/jeamon/digital-root-algorithms       2.500s

// all tests and benchmarks results

[13:14:34] {nxos-geek}:~$ go test -run=. -bench=. -benchmem -benchtime=20000000x -v
=== RUN   TestDigitalRootWithString
--- PASS: TestDigitalRootWithString (0.00s)
=== RUN   TestDigitalRootWithModulo
--- PASS: TestDigitalRootWithModulo (0.00s)
goos: windows
goarch: amd64
pkg: github.com/jeamon/digital-root-algorithms
BenchmarkDigitalRootWithString-4        20000000                98.1 ns/op             4 B/op          1 allocs/op
BenchmarkDigitalRootWithModulo-4        20000000                13.1 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/jeamon/digital-root-algorithms       2.294s

*/
