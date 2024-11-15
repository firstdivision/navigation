package navigation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLatLon(t *testing.T) {
	var tests = []struct {
		name string
		lat  float64
		lon  float64
	}{
		{
			"Ok 1",
			11.22,
			33.44,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			latLon := NewLatLon(tc.lat, tc.lon)
			assert.Equal(t, tc.lat, latLon.Latitude())
			assert.Equal(t, tc.lon, latLon.Longitude())
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
		{
			"Rio de Janeiro Brazil to Bangkok Thailand",
			LatLon{latitude: 22.55, longitude: 43.12},
			LatLon{latitude: 13.45, longitude: 100.28},
			6094.544408786774,
		},
		{
			"Identical locations",
			LatLon{latitude: 22.55, longitude: 43.12},
			LatLon{latitude: 22.55, longitude: 43.12},
			0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			output := tc.p1.DistanceTo(tc.p2)
			assert.Equal(t, output, tc.distance)
		})
	}
}
