package openweathermap

type (
	ForecastWather struct {
		Cod     string
		Message float64
		City    City
		Cnt     int64
		List    []CurrentWeather
	}

	CurrentWeather struct {
		Coord   Coord
		Wind    Wind
		Main    Main
		Clouds  Clouds
		Weather []Weather
		Sys     Sys
		Name    string
		Base    string
		Dt      int64
		Id      int64
		Cod     int64
	}

	City struct {
		Id         int64
		Name       string
		Coord      Coord
		Country    string
		Population int64
	}

	Weather struct {
		Id          int64
		Main        string
		Description string
	}

	Clouds struct {
		all int64
	}

	Main struct {
		Temp     float64
		Pressure float64
		Humidity float64
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
		Deg   float64
	}

	Coord struct {
		Lon float64
		Lat float64
	}
)
