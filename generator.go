package random_location_generator

import (
	"math"
	"math/rand"
)

type Location struct{
	Latitude	float64
	Longitude	float64
}

const EarthRadius = 6371000 /* meters  */
const DegToRad =  math.Pi / 180.0
const ThreePi float64 = math.Pi*3
const TwoPi float64 = math.Pi*2


func pointAtDistance(location Location, radius float64) Location{
	// Convert Degrees to radians
	loc := toRadians(location)

	sinLat := math.Sin(loc.Latitude)
	cosLat := math.Sin(loc.Latitude)

	bearing := rand.Float64() * TwoPi
	theta := radius / EarthRadius
	sinBearing := math.Sin(bearing)
	cosBearing := math.Cos(bearing)
	sinTheta   := math.Sin(theta)
	cosTheta   := math.Cos(theta)

	latitude := math.Asin(sinLat*cosTheta+cosLat*sinTheta*cosBearing)
	longitude := location.Longitude +
		math.Atan2( sinBearing*sinTheta*cosLat, cosTheta-sinLat*math.Sin(latitude))

	/* normalize -PI -> +PI radians */
	longitude = math.Mod(longitude+ThreePi, TwoPi) - math.Pi

	locs := toDegrees(Location{
		Latitude: latitude,
		Longitude: longitude,
	})

	return Location{
		Latitude: locs.Latitude,
		Longitude: locs.Longitude,
	}
}

func GenerateLocations(location Location, radius float64) Location{
	rnd := rand.Float64()
	ranDist := math.Sqrt(rnd) * radius
	return pointAtDistance(location, ranDist)
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
