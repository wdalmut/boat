package gps

import(
    // #cgo LDFLAGS: -lgps -lm
    // #include <gps.h>
    "C"
    "math"
)

const (
    Earth = 6371
    DegToRad = 180/math.Pi
    RadToDeg = math.Pi/180

)

func GpsLocation() (float64, float64) {
    var data C.loc_t
    var lat, lon float64

    C.gps_location(&data)

    lat = float64(data.latitude)
    lon = float64(data.longitude)

    return lat, lon
}

func Distance(lat1, lon1, lat2, lon2 float64) float64 {
    dLat := (lat2 - lat1) * DegToRad
    dLon := (lon2 - lon1) * DegToRad

    lat1 = lat1 * DegToRad
    lat2 = lat2 * DegToRad

    a := math.Sin(dLat/2) * math.Sin(dLat/2) + math.Sin(dLon/2) * math.Sin(dLon/2) * math.Cos(lat1) * math.Cos(lat2)
    c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a));

    d := Earth * c;
    return d;
}
