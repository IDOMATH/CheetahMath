package formulas

import "math"

// Haversine accepts a latitude and longitude pairs as float64s
// and returns the distance between them in km.
func Haversine(lat1, long1, lat2, long2 float64) float64 {
	dLat := ToRadians(lat2 - lat1)
	dLong := ToRadians(long2 - long1)

	rLat1 := ToRadians(lat1)
	rlat2 := ToRadians(lat2)

	a := math.Pow(math.Sin(dLat/2.0), 2.0) +
		math.Pow(math.Sin(dLong/2.0), 2.0)*
			math.Cos(rLat1)*math.Cos(rlat2)

	radius := 6371000.0

	c := 2.0 * math.Asin(math.Sqrt(a))
	return radius * c
}
