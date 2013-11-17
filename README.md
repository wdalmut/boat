# GPS golang library

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
BenchmarkMidpoint        5000000       401 ns/op
BenchmarkDistance       10000000       202 ns/op
BenchmarkDistance2      10000000       180 ns/op
ok      github.com/wdalmut/boat/gps  6.664s
```

