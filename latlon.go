package navigation

import "math"

// LatLon structure represents a coordinate on the surface of the Earth
type LatLon struct {
	Latitude  float64
	Longitude float64
}

// DistanceTo returns the distance, in kilometers, between the two points
func (l *LatLon) DistanceTo(point LatLon) float64 {

	lat1 := toRadians(l.Latitude)
	lon1 := toRadians(l.Longitude)

	lat2 := toRadians(point.Latitude)
	lon2 := toRadians(point.Longitude)

	dLat := lat2 - lat1
	dLon := lon2 - lon1

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	var d = EarthRadius * c

	return d
}

// toRadians converts numeric degrees to radians
func toRadians(f float64) float64 {
	return f * math.Pi / 180
}
