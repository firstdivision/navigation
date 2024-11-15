package navigation

import "math"

type ILatLon interface {
	DistanceTo(point LatLon) float64
	Latitude() float64
	Longitude() float64
}

// LatLon structure represents a coordinate on the surface of the Earth
type LatLon struct {
	latitude  float64
	longitude float64
}

func NewLatLon(latitude, Longitude float64) ILatLon {
	return &LatLon{
		latitude:  latitude,
		longitude: Longitude,
	}
}

func (l *LatLon) Latitude() float64 {
	return l.latitude
}

func (l *LatLon) Longitude() float64 {
	return l.longitude
}

func (l *LatLon) DistanceTo(point LatLon) float64 {

	lat1 := toRadians(l.latitude)
	lon1 := toRadians(l.longitude)

	lat2 := toRadians(point.latitude)
	lon2 := toRadians(point.longitude)

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

// // toDegrees converts numeric degrees to degrees
// func toDegrees(f float64) float64 {
// 	return f * 180 / math.Pi
// }
