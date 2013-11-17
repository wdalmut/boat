package gps

import(
    // #cgo LDFLAGS: -lgps -lm
    // #include <gps.h>
    "C"
    "math"
)

const (
    Earth = 6356.7523
    DegToRad = math.Pi/180
    RadToDeg = 180/math.Pi

)

func GpsLocation() (float64, float64) {
    var data C.loc_t
    var lat, lon float64

    C.gps_location(&data)

    lat = float64(data.latitude)
    lon = float64(data.longitude)

    return lat, lon
}

func Midpoint(lat1, lon1, lat2, lon2 float64) (float64, float64) {
    dLon := (lon2 - lon1) * DegToRad

    lat1 = lat1 * DegToRad
    lat2 = lat2 * DegToRad

    bx := math.Cos(lat2) * math.Cos(dLon)
    by := math.Cos(lat2) * math.Sin(dLon)

    lat := math.Atan2(math.Sin(lat1) + math.Sin(lat2), math.Sqrt(
        math.Pow(math.Cos(lat1)+bx, 2) + math.Pow(by, 2))) * RadToDeg

    lon := lon1 + math.Atan2(by, math.Cos(lat1) + bx) * RadToDeg

    return lat, lon
}

// Spherical Law of Cosines
func Distance2(lat1, lon1, lat2, lon2 float64) float64 {
    return math.Acos(
        math.Sin(lat1 * DegToRad)*math.Sin(lat2 * DegToRad) +
        math.Cos(lat1 * DegToRad) * math.Cos(lat2 * DegToRad) *
        math.Cos(lon2 *DegToRad - lon1 * DegToRad)) * Earth;
}

// Haversine formula
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
    dLat := (lat2 - lat1) * DegToRad
    dLon := (lon2 - lon1) * DegToRad

    lat1 = lat1 * DegToRad
    lat2 = lat2 * DegToRad

    a1 := math.Sin(dLat/2) * math.Sin(dLat/2)
    a2 := math.Sin(dLon/2) * math.Sin(dLon/2) * math.Cos(lat1) * math.Cos(lat2)
    a := a1 + a2
    c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a));

    return Earth * c;
}
