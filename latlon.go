package navigation

import (
	"fmt"
	"math"
)

// latLon structure represents a coordinate on the surface of the Earth
type latLon struct {
	Latitude  float64
	Longitude float64
}

func NewLatLon(lat, lon float64) (*latLon, error) {

	if lat < -90 || lat > 90 {
		return nil, fmt.Errorf("latitude must be between -90 and 90")
	}

	if lon < -180 || lon > 180 {
		return nil, fmt.Errorf("longitude must be between -180 and 180")
	}

	return &latLon{
		Latitude:  lat,
		Longitude: lon,
	}, nil
}

// DistanceTo returns the distance, in kilometers, between the two points
func (l *latLon) DistanceTo(point latLon) float64 {

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
