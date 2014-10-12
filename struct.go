package openweathermap

type (
	Current struct {
		Coord   Coord
		Wind    Wind
		Main    Main
		Clouds  Clouds
		Weather []WeatherStruct
		Sys     Sys
		Name    string
		Base    string
		Dt      int64
		Id      int64
		Cod     int64
	}

	WeatherStruct struct {
		Id          int64
		Main        string
		Description string
	}

	Clouds struct {
		all int64
	}

	Main struct {
		Temp     float64
		Pressure int64
		Humidity int64
		Temp_min float64
		Temp_max float64
	}

	Sys struct {
		Type    int64
		Id      int64
		Country string
		Message float64
		Sunset  int64
		Sundown int64
	}

	Wind struct {
		Speed float64
		Deg   int64
	}

	Coord struct {
		Lon float64
		Lat float64
	}
)
