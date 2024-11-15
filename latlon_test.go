package navigation

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLatLon(t *testing.T) {
	var tests = []struct {
		name string
		lat  float64
		lon  float64

		//expected
		latLon *LatLon
		err    error
	}{
		{
			name:   "Ok 1 - zeros",
			lat:    0,
			lon:    0,
			latLon: &LatLon{},
			err:    nil,
		}, {
			name:   "Bad 1 - Latitude below 90",
			lat:    -91,
			lon:    0,
			latLon: nil,
			err:    errors.New("latitude must be between -90 and 90"),
		}, {
			name:   "Bad 2 - Latitude above 90",
			lat:    91,
			lon:    0,
			latLon: nil,
			err:    errors.New("latitude must be between -90 and 90"),
		}, {
			name:   "Bad 3 - Longitude below -360",
			lat:    0,
			lon:    -361,
			latLon: nil,
			err:    errors.New("longitude must be between -180 and 180"),
		}, {
			name:   "Bad 4 - Longitude above 360",
			lat:    0,
			lon:    361,
			latLon: nil,
			err:    errors.New("longitude must be between -180 and 180"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			latLon, err := NewLatLon(tc.lat, tc.lon)

			assert.Equal(t, tc.latLon, latLon)
			assert.Equal(t, tc.err, err)
		})
	}
}

func TestDistanceTo(t *testing.T) {

	var tests = []struct {
		name     string
		p1       LatLon
		p2       LatLon
		distance float64
	}{
		// Test case borrowed from https://github.com/umahmood/haversine
		{
			"Rio de Janeiro Brazil to Bangkok Thailand",
			LatLon{Latitude: 22.55, Longitude: 43.12},
			LatLon{Latitude: 13.45, Longitude: 100.28},
			6094.544408786774,
		},
		{
			"NYC to Boston",
			LatLon{Latitude: 40.71, Longitude: 74},
			LatLon{Latitude: 42.36, Longitude: 71.05},
			306.4858388776133,
		},
		{
			"Identical locations",
			LatLon{Latitude: 22.55, Longitude: 43.12},
			LatLon{Latitude: 22.55, Longitude: 43.12},
			0,
		},
		{
			"Max Values",
			LatLon{Latitude: -90, Longitude: -360},
			LatLon{Latitude: 90, Longitude: 360},
			20015.086796020572,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			output := tc.p1.DistanceTo(tc.p2)
			assert.Equal(t, output, tc.distance)
		})
	}
}
