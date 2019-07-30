package mathutil

import "testing"

func TestDistance(t *testing.T) {
    point1 := []float64{0, 0, 0}
    point2 := []float64{1, 1, 1}
    point3 := []float64{-1, -2, -3}
    point4 := []float64{1, 2}

    // Test #1: Distance between a point and itself should be 0
    if distance, _ := SquaredDistance(point1, point1); distance != 0.0 {
        t.Error("Distance between a point and itself is not 0")
    }

    // Test #2: Distance between two points should be symmetric
    distance1, _ := SquaredDistance(point1, point2)
    distance2, _ := SquaredDistance(point2, point1)

    if distance1 != distance2 {
        t.Error("Distance function is not symmetric")
    }

    // Test #3: Distance should obey triangle inequality
    distance3, _ := SquaredDistance(point1, point2)
    distance4, _ := SquaredDistance(point2, point3)
    distance5, _ := SquaredDistance(point1, point3)

    if distance5 > (distance3 + distance4) {
        t.Error("Distance function does not obey triangle inequality")
    }

    // Test #4: Distance function should produce an error if points have different number of dimensions
    if _, err := SquaredDistance(point1, point4); err == nil {
        t.Error("Distance function doesn't produce an error when points have different number of dimensions")
    }
}