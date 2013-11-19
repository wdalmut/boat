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
}


func TestPrepareMidpoints(t *testing.T) {

    for _, tt := range tests {
        m := new(Mission)
        for j:=1;j<len(tt.points);j++ {
            /*distance := geo.Distance(&tt.points[j-1], &tt.points[j])*/
            m.AddGoalPoint(&tt.points[j-1])
            m.AddGoalPoint(&tt.points[j])
        }

        m.GenerateCheckpoints()

        for j:=0;j<len(tt.checkpoints);j++ {
            if tt.checkpoints[j] != len(m.checkpoints[j]) {
                t.Errorf("Invalid checkpoints count %v, want %v", len(m.checkpoints[j]), tt.checkpoints[j])
            }
        }
    }
}

