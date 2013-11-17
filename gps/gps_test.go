package gps

import(
    "testing"
    "math"
)

func TestDistance(t *testing.T) {
    lat1 := 50.066389
    lon1 := 5.714722

    lat2 := 58.643889
    lon2 := 3.07

    want := 966.68

    if distance := Distance(lat1, lon1, lat2, lon2); math.Abs(distance - want) > 0.01 {
        t.Errorf("Wrong distance: %v, want %v", distance, want)
    }
}

func TestDistance2(t *testing.T) {
    lat1 := 50.066389
    lon1 := 5.714722

    lat2 := 58.643889
    lon2 := 3.07

    want := 966.68

    if distance := Distance2(lat1, lon1, lat2, lon2); math.Abs(distance - want) > 0.01 {
        t.Errorf("Wrong distance: %v, want %v", distance, want)
    }
}

func TestMidpoint(t *testing.T) {
    lat1 := 50.066389
    lon1 := 5.714722

    lat2 := 58.643889
    lon2 := 3.07

    lat, lon := Midpoint(lat1, lon1, lat2, lon2)

    wlat := 54.36
    wlon := 4.53

    if math.Abs(lat - wlat) > 0.01 {
        t.Errorf("Wrong distance: %v, want %v", lat, wlat)
    }

    if math.Abs(lon - wlon) > 0.01 {
        t.Errorf("Wrong distance: %v, want %v", lon, wlon)
    }
}
