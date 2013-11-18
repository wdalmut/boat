# GPS golang library

[![Build Status](https://travis-ci.org/wdalmut/boat.png?branch=develop)](https://travis-ci.org/wdalmut/boat)

The GPS library (mixed C and Go)

## Build it

You need this C library: https://github.com/wdalmut/libgps

```shell
$ git clone https://github.com/wdalmut/libgps.git
$ cd libgps
$ make
$ sudo make install # Sym links into /usr/lib and /usr/include
```

## Run it

```shell
$ go run main/eg.go
```

## Tests & Benchmarks

Run tests only

```shell
go test
```

Tests + Benchmarks

```shell
go test -test.bench .
```

```
PASS
BenchmarkMidpoint	 5000000	       472 ns/op
BenchmarkDistance	10000000	       204 ns/op
BenchmarkDistance2	10000000	       178 ns/op
BenchmarkDistance3	10000000	       154 ns/op
ok  	github.com/wdalmut/boat/gps	8.777s
```

## Raspberry Pi Bench

```
PASS
BenchmarkMidpoint	  100000	     16270 ns/op
BenchmarkDistance	  200000	      8039 ns/op
BenchmarkDistance2	  500000	      6543 ns/op
BenchmarkDistance3	  500000	      6005 ns/op
ok  	github.com/wdalmut/boat/gps	9.987s
```

