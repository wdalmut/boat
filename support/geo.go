package support

import (
    "math"
    "github.com/wdalmut/boat/geo"
)

func Midpoint(lat1, lon1, lat2, lon2 float64) (float64, float64) {
    dLon := (lon2 - lon1) * geo.DegToRad

    lat1 = lat1 * geo.DegToRad
    lat2 = lat2 * geo.DegToRad

    bx := math.Cos(lat2) * math.Cos(dLon)
    by := math.Cos(lat2) * math.Sin(dLon)

    lat := math.Atan2(math.Sin(lat1) + math.Sin(lat2), math.Sqrt(
        math.Pow(math.Cos(lat1)+bx, 2) + math.Pow(by, 2))) * geo.RadToDeg

    lon := lon1 + math.Atan2(by, math.Cos(lat1) + bx) * geo.RadToDeg

    return lat, lon
}

// Pythagora's theorem
func Distance3(lat1, lon1, lat2, lon2 float64) float64 {
    dlon := (lon2 - lon1) * geo.DegToRad
    lat := (lat1 + lat2) * geo.DegToRad
    dlat := (lat2 -lat1) * geo.DegToRad

    x := dlon * math.Cos(lat/2)

    return math.Sqrt(math.Pow(x, 2) + math.Pow(dlat, 2)) * geo.Earth
}

// Spherical Law of Cosines
func Distance2(lat1, lon1, lat2, lon2 float64) float64 {
    return math.Acos(
        math.Sin(lat1 * geo.DegToRad)*math.Sin(lat2 * geo.DegToRad) +
        math.Cos(lat1 * geo.DegToRad) * math.Cos(lat2 * geo.DegToRad) *
        math.Cos(lon2 *geo.DegToRad - lon1 * geo.DegToRad)) * geo.Earth
}

// Haversine formula
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
    dLat := (lat2 - lat1) * geo.DegToRad
    dLon := (lon2 - lon1) * geo.DegToRad

    lat1 = lat1 * geo.DegToRad
    lat2 = lat2 * geo.DegToRad

    a1 := math.Sin(dLat/2) * math.Sin(dLat/2)
    a2 := math.Sin(dLon/2) * math.Sin(dLon/2) * math.Cos(lat1) * math.Cos(lat2)
    a := a1 + a2
    c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a));

    return geo.Earth * c;
}

