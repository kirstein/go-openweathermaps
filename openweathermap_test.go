package openweathermap

import (
	"log"
	"testing"
)

func TestCurrentCoords(t *testing.T) {
	q, err := Current("tallinn", "metric")
	if err != nil {
		log.Panic(err)
	}

	if q.Coord.Lat != 59.44 {
		t.Errorf("Invalid latitude = %f, expected %2f", q.Coord.Lat, 59.44)
	}
	if q.Coord.Lon != 24.75 {
		t.Errorf("Invalid longitude = %f, expected %2f", q.Coord.Lon, 24.75)
	}
}

func TestCurrentWind(t *testing.T) {
	q, err := Current("tallinn", "metric")
	if err != nil {
		log.Panic(err)
	}

	if q.Wind.Speed == 0 {
		t.Errorf("Wind speed should be set")
	}
	if q.Wind.Deg == 0 {
		t.Errorf("Wind deg should be set")
	}
}

func TestForecast(t *testing.T) {
	q, err := Forecast("tallinn", "metric")
	if err != nil {
		log.Panic(err)
	}
	if q.City.Name != "Tallinn" {
		t.Errorf("City = %s, expected %s", q.City.Name, "tallinn")
	}
}
