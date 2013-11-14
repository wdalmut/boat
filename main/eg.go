package main

import(
    "github.com/wdalmut/boat/gps"
    "fmt"
)

func main() {
    lat, lon := gps.GpsLocation()

    fmt.Printf("%v, %v\n", lat, lon)
}
