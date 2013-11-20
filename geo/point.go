package geo

import (
    "math"
)

const (
    Earth = 6356.7523
    DegToRad = math.Pi/180
    RadToDeg = 180/math.Pi

)

type Point struct {
    Latitude float64
    Longitude float64
}

func Midpoint(point1, point2 *Point) *Point {
    midpoint := new(Point)

    dLon := (point2.Longitude - point1.Longitude) * DegToRad

    lat1 := point1.Latitude * DegToRad
    lat2 := point2.Latitude * DegToRad

    bx := math.Cos(lat2) * math.Cos(dLon)
    by := math.Cos(lat2) * math.Sin(dLon)

    midpoint.Latitude = math.Atan2(math.Sin(lat1) + math.Sin(lat2), math.Sqrt(
        math.Pow(math.Cos(lat1)+bx, 2) + math.Pow(by, 2))) * RadToDeg

    midpoint.Longitude = point1.Longitude + math.Atan2(by, math.Cos(lat1) + bx) * RadToDeg

    return midpoint
}

// Calculate distance using spherical law of cosine
func Distance(point1, point2 *Point) float64 {
    return math.Acos(
        math.Sin(point1.Latitude * DegToRad)*math.Sin(point2.Latitude * DegToRad) +
        math.Cos(point1.Latitude * DegToRad) * math.Cos(point2.Latitude * DegToRad) *
        math.Cos(point2.Longitude *DegToRad - point1.Longitude * DegToRad)) * Earth
}

// Calculate the angle between 2 points
func Bearing(point1, point2 *Point) float64 {
    dLon := (point2.Longitude - point1.Longitude) * DegToRad
    lat1 := point1.Latitude * DegToRad
    lat2 := point2.Latitude * DegToRad

    y := math.Sin(dLon) * math.Cos(lat2);
    x := math.Cos(lat1) * math.Sin(lat2) - math.Sin(lat1) * math.Cos(lat2)*math.Cos(dLon)

    return math.Atan2(y, x) * RadToDeg
}

