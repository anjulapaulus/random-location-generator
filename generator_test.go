package random_location_generator

import (
	"testing"
)

func TestGetLocation(t *testing.T) {
	location:= GenerateLocations(Location{
		Latitude: 6.981106276681917,
		Longitude: 79.87749937315392,
	}, 100)
	t.Error(location)
}
