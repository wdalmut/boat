package mission

import (
    "fmt"
    "github.com/wdalmut/boat/geo"
)

// Configuration
const (
    // 1 km max between points
    MAX_DISTANCE = 1
)

// The mission
// TODO checkpoints mem optimization
type Mission struct {
    goals []*geo.Point
    checkpoints [][]geo.Point
}

// Add goal points
func (m *Mission)AddGoalPoint(point *geo.Point) {
    m.goals = append(m.goals, point)
}

// Generate intermediate checkpoints
func (m *Mission)GenerateCheckpoints() {
    if (len(m.goals) <= 1) {
        panic("We need more goals points in order to calculate checkpoints")
    }

    m.checkpoints = make([][]geo.Point, (len(m.goals)-1))

    for i:=1; i< len(m.goals); i++ {
        m.prepareMidpoints(i-1, m.goals[i-1], m.goals[i])
    }
}

// Calculate all checkpoints (in order)
func (m *Mission)prepareMidpoints(slice int, first, second *geo.Point) {
    distance := geo.Distance(first, second)

    if distance < MAX_DISTANCE {
        return
    }

    midpoint := geo.Midpoint(first, second)

    m.prepareMidpoints(slice, first, midpoint)
    m.checkpoints[slice] = append(m.checkpoints[slice], *midpoint)
    m.prepareMidpoints(slice, midpoint, second)
}

// Help debug printing CSV list of checkpoints
// TODO remove this method
func (m *Mission)PrintCsv() {
    for i:=0; i<len(m.checkpoints); i++ {
        fmt.Printf("Mission %v,_,_,_\n", i)
        for j:=0; j<len(m.checkpoints[i]); j++ {
            fmt.Printf("%v, %v, %v, %v\n", m.checkpoints[i][j].Latitude, m.checkpoints[i][j].Longitude, i, j)
        }
    }
}

