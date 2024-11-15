package navigation

import "math"

type ILatLon interface {
	DistanceTo(point LatLon, precision int) float64
}

// LatLon structure represents a coordinate on the surfact of the Earth
type LatLon struct {
	Latitude  float64
	Longitude float64
}

func NewLatLon(latitude, Longitude float64) ILatLon {
	return &LatLon{
		Latitude:  latitude,
		Longitude: Longitude,
	}
}

func (l *LatLon) DistanceTo(point LatLon, precision int) float64 {

	if precision <= 0 {
		precision = 4
	}

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

// toRadians converts numeric degrees to degrees
func toDegrees(f float64) float64 {
	return f * 180 / math.Pi
}
