package random_location_generator

import (
	"math"
	"math/rand"
)
const EarthRadius float64= 6372796.924
const DegToRad =  math.Pi / 180.0

type Location struct{
	Latitude	float64
	Longitude	float64
}

func RandomLocation(location Location, radius float64) Location{
	rad := toRadians(location)

	rand1 := rand.Float64()
	rand2 := rand.Float64()

	maxdist:= radius / EarthRadius

	dist := math.Acos(rand1*(math.Cos(maxdist) - 1) + 1)
	brg := 2*math.Pi*rand2

	lat := math.Asin(math.Sin(rad.Latitude)*math.Cos(dist) + math.Cos(rad.Latitude)*math.Sin(dist)*math.Cos(brg))
	lon := rad.Longitude + math.Atan2(math.Sin(brg)*math.Sin(dist)*math.Cos(rad.Latitude), math.Cos(dist)-math.Sin(rad.Latitude)*math.Sin(lat))

	lon = lon + 2*math.Pi

	if lon < -math.Pi{
		lon = lon + 2*math.Pi
	}else if lon > math.Pi{
		lon = lon - 2*math.Pi
	}

	loc:=toDegrees(Location{
		Latitude: lat,
		Longitude: lon,
	})
	return loc

}

func toRadians(location Location) Location{
	lat := location.Latitude * DegToRad
	lon := location.Longitude * DegToRad

	return Location{
		Latitude: lat,
		Longitude: lon,
	}
}

func toDegrees(location Location) Location{
	lat := location.Latitude / DegToRad
	lon := location.Longitude / DegToRad

	return Location{
		Latitude: lat,
		Longitude: lon,
	}
}