package main

import(
    "fmt"
    "github.com/wdalmut/boat/gps"
    "github.com/wdalmut/boat/mission"
    "github.com/wdalmut/boat/geo"
)

// An example of mission explosion in checkpoints from goals
func main() {
    lat, lon := gps.GpsLocation()

    fmt.Printf("%v, %v\n", lat, lon)


    m := new(mission.Mission)

    point1 := new(geo.Point)
    point1.Latitude = 43.91
    point1.Longitude = 8.1125

    point2 := new(geo.Point)
    point2.Latitude = 44.055556
    point2.Longitude = 9.816944

    m.AddGoalPoint(point1)
    m.AddGoalPoint(point2)

    m.GenerateCheckpoints()

    m.PrintCsv()
}
