package navigation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistanceTo(t *testing.T) {

	var tests = []struct {
		name     string
		p1       LatLon
		p2       LatLon
		distance float64
	}{
		{
			"Rio de Janeiro Brazil to Bangkok Thailand",
			LatLon{Latitude: 22.55, Longitude: 43.12},
			LatLon{Latitude: 13.45, Longitude: 100.28},
			6094.544408786774,
		},
		{
			"Identical locations",
			LatLon{Latitude: 22.55, Longitude: 43.12},
			LatLon{Latitude: 22.55, Longitude: 43.12},
			0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			output := tc.p1.DistanceTo(tc.p2, 0)
			assert.Equal(t, output, tc.distance)
		})
	}
}
