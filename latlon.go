package navigation

import (
	"fmt"
	"math"
)

// LatLon structure represents a coordinate on the surface of the Earth
type LatLon struct {
	Latitude  float64
	Longitude float64
}

func NewLatLon(lat, lon float64) (*LatLon, error) {

	if lat < -90 || lat > 90 {
		return nil, fmt.Errorf("latitude must be between -90 and 90")
	}

	if lon < -180 || lon > 180 {
		return nil, fmt.Errorf("longitude must be between -180 and 180")
	}

	return &LatLon{
		Latitude:  lat,
		Longitude: lon,
	}, nil
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

// BearingTo returns the (initial) bearing from this point to the supplied point, in degrees

/*
JavaScript:
LatLon.prototype.finalBearingTo = function (point) {
    // get initial bearing from supplied point back to this point...
    var lat1 = point._lat.toRad()
	var lat2 = this._lat.toRad();
    var dLon = (this._lon - point._lon).toRad();

    var y = Math.sin(dLon) * Math.cos(lat2);
    var x = Math.cos(lat1) * Math.sin(lat2) -
          Math.sin(lat1) * Math.cos(lat2) * Math.cos(dLon);
    var brng = Math.atan2(y, x);

    // ... & reverse it by adding 180Â°
    return (brng.toDeg() + 180) % 360;
}
*/

// func (this *LatLon) BearingTo(point LatLon) float64 {
// 	lat1 := toRadians(point.Latitude)
// 	lat2 := toRadians(this.Latitude)
// 	dLon := toRadians(this.Longitude - point.Longitude)

// 	y := math.Sin(dLon) * math.Cos(lat2)
// 	x := math.Cos(lat1)*math.Sin(lat2) -
// 		math.Sin(lat1)*math.Cos(lat2)*math.Cos(dLon)
// 	brng := math.Atan2(y, x)
// 	brngDegrees := toDegrees(brng)

// 	output := math.Mod(brngDegrees+180, 360)

// 	return output
// }

// toRadians converts numeric degrees to radians
func toRadians(f float64) float64 {
	return f * math.Pi / 180
}

// toDegrees converts radians to numeric (signed) degrees
func toDegrees(f float64) float64 {
	return f * 180 / math.Pi
}
