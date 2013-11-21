package mission

import (
    "testing"
    "github.com/wdalmut/boat/geo"
)

var tests = []struct{
    points []geo.Point
    checkpoints []int
} {
    {[]geo.Point{geo.Point{50.066389, 5.714722}, geo.Point{58.643889, 3.07}}, []int{1023}},
    {[]geo.Point{geo.Point{50.066389, 5.714722}, geo.Point{50.166389, 5.714723}}, []int{15}},
    {[]geo.Point{geo.Point{43.91, 8.1125}, geo.Point{44.055556, 9.816944}, geo.Point{43.54750, 10.28032}}, []int{255, 127}},
}

// Test not valid mission (0 points)
func TestNoPoints(t *testing.T) {
    defer func() {
        s := recover()
        if s == nil {
            t.Errorf("did not panic...")
        }
    }()

    m := new(Mission)
    m.GenerateCheckpoints()
}

// Test not valid mission (one point only)
func TestOnePoint(t *testing.T) {
    defer func() {
        s := recover()
        if s == nil {
            t.Errorf("did not panic...")
        }
    }()

    m := new(Mission)
    m.AddGoalPoint(new(geo.Point))
    m.GenerateCheckpoints()
}

// Test midpoints generation
func TestPrepareMidpoints(t *testing.T) {

    for _, tt := range tests {
        m := new(Mission)
        for j:=0;j<len(tt.points);j++ {
            m.AddGoalPoint(&tt.points[j])
        }

        m.GenerateCheckpoints()

        if steps := len(m.checkpoints); steps != (len(tt.points)-1) {
            t.Errorf("Invalid number of steps: %v, want %v", steps, len(tt.points)-1)
        }

        for j:=0;j<len(tt.checkpoints);j++ {
            if tt.checkpoints[j] != len(m.checkpoints[j]) {
                t.Errorf("Invalid checkpoints count %v, want %v", len(m.checkpoints[j]), tt.checkpoints[j])
            }
        }
    }
}

// Benchmark midpoints generation
func BenchmarkGenerateCheckpoints(b *testing.B) {
    m := new(Mission)

    point1 := new(geo.Point)
    point1.Latitude = 50.066389
    point1.Longitude = 5.714722

    point2 := new(geo.Point)
    point2.Latitude = 58.643889
    point2.Longitude = 3.07

    m.AddGoalPoint(point1)
    m.AddGoalPoint(point2)

    for i:=0; i<b.N; i++ {
        m.GenerateCheckpoints();
    }
}

