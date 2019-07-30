package mathutil

import (
    "errors"
    "math"
)

// Calculate squared Euclidean distance between two points
func SquaredDistance(point1 []float64, point2 []float64) (float64, error) {
    if len(point1) != len(point2) {
        return 0, errors.New("Points do not have the same number of dimensions")
    }

    result := 0.0

    for i := 0; i < len(point1); i++ {
        if point1[i] < point2[i] {
            result += math.Pow(point2[i] - point1[i], 2)
        } else {
            result += math.Pow(point1[i] - point2[i], 2)
        }
    }

    return result, nil
}