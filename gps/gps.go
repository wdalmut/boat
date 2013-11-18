package gps

import(
    // #cgo LDFLAGS: -lgps -lm
    // #include <gps.h>
    "C"
)

func GpsLocation() (float64, float64) {
    var data C.loc_t
    var lat, lon float64

    C.gps_location(&data)

    lat = float64(data.latitude)
    lon = float64(data.longitude)

    return lat, lon
}

