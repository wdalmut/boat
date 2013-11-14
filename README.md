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

