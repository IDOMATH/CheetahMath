package haversine

import (
	"github.com/IDOMATH/CheetahMath/radians"
	"math"
)

// Haversine accepts a latitude and longitude pairs as float64s
// and returns the distance between them in radians.
// Multiply the result by the radius of the earth.
func Haversine(lat1, long1, lat2, long2 float64) float64 {
	dLat := radians.ToRadians(lat2 - lat1)
	dLong := radians.ToRadians(long2 - long1)

	rLat1 := radians.ToRadians(lat1)
	rlat2 := radians.ToRadians(lat2)

	a := math.Pow(math.Sin(dLat/2.0), 2.0) +
		math.Pow(math.Sin(dLong/2.0), 2.0)*
			math.Cos(rLat1)*math.Cos(rlat2)

	return 2.0 * math.Asin(math.Sqrt(a))
}
