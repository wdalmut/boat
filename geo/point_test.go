package geo

import (
    "testing"
    "math"
)

func TestMidpoint(t *testing.T) {
    point1 := new(Point)
    point1.Latitude = 50.066389
    point1.Longitude = 5.714722

    point2 := new(Point)
    point2.Latitude = 58.643889
    point2.Longitude = 3.07

    midpoint := Midpoint(point1, point2)

    wlat := 54.36
    wlon := 4.53

    if math.Abs(midpoint.Latitude - wlat) > 0.01 {
        t.Errorf("Wrong distance: %v, want %v", midpoint.Latitude, wlat)
    }

    if math.Abs(midpoint.Longitude - wlon) > 0.01 {
        t.Errorf("Wrong distance: %v, want %v", midpoint.Longitude, wlon)
    }
}

func TestDistance(t *testing.T) {
    point1 := new(Point)
    point1.Latitude = 50.066389
    point1.Longitude = 5.714722

    point2 := new(Point)
    point2.Latitude = 58.643889
    point2.Longitude = 3.07

    want := 966.68

    if distance := Distance(point1, point2); math.Abs(distance - want) > 0.01 {
        t.Errorf("Wrong distance: %v, want %v", distance, want)
    }
}

func BenchmarkMidpoint(b *testing.B) {
    point1 := new(Point)
    point1.Latitude = 50.066389
    point1.Longitude = 5.714722

    point2 := new(Point)
    point2.Latitude = 58.643889
    point2.Longitude = 3.07

    for i := 0; i < b.N; i++ {
        Midpoint(point1, point2)
    }
}

func BenchmarkDistance(b *testing.B) {
    point1 := new(Point)
    point1.Latitude = 50.066389
    point1.Longitude = 5.714722

    point2 := new(Point)
    point2.Latitude = 58.643889
    point2.Longitude = 3.07

    for i := 0; i < b.N; i++ {
        Distance(point1, point2)
    }
}

