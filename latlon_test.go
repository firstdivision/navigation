package navigation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	p1       LatLon
	p2       LatLon
	distance float64
}{
	{
		LatLon{Latitude: 22.55, Longitude: 43.12},
		LatLon{Latitude: 13.45, Longitude: 100.28},
		6094.544408786774,
	},
}

func TestDistanceTo(t *testing.T) {
	for _, tc := range tests {
		output := tc.p1.DistanceTo(tc.p2, 0)
		assert.Equal(t, output, tc.distance)
	}
}
