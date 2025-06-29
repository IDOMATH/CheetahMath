package radians

import "math"

func ToRadians(degrees float64) float64 {
	return degrees / 180.0 * math.Pi
}
