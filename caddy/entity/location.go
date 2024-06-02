package entity

import (
	"math"
	"math/rand/v2"
)

type LocationZone struct {
	StartPoint     LocationPoint
	EndPoint       LocationPoint
	LocationPoints []LocationPoint
}

type LocationPoint struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

func GenerateFlatSquareLocationBounds(startingPoint LocationPoint, area float64) (float64, float64, float64, float64) {
	sideLength := math.Sqrt(area)

	x1 := startingPoint.Lat
	y1 := startingPoint.Long

	x2 := x1 + sideLength
	y2 := y1 + sideLength

	return x1, y1, x2, y2
}

func GenerateRandomLocation(startingPoint, endPoint LocationPoint) LocationPoint {
	latitude := rand.Float64()*(endPoint.Lat-startingPoint.Lat) + startingPoint.Lat
	longitude := rand.Float64()*(endPoint.Long-startingPoint.Long) + startingPoint.Long

	return LocationPoint{
		Lat:  latitude,
		Long: longitude,
	}
}

func CalculateTimeInMinute(distanceInKm float64, speedInKmHour int) int {
	return int(distanceInKm / float64(speedInKmHour) * 60)
}

func CalculateDistance(p1, p2 LocationPoint) float64 {
	const R = 6371 // Radius of Earth in kilometers

	lat1 := p1.Lat * math.Pi / 180
	long1 := p1.Long * math.Pi / 180
	lat2 := p2.Lat * math.Pi / 180
	long2 := p2.Long * math.Pi / 180

	dlat := lat2 - lat1
	dlong := long2 - long1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dlong/2)*math.Sin(dlong/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}
