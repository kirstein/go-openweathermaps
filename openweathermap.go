package openweathermap

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	version     = "2.5"
	baseUrl     = "http://api.openweathermap.org/data/" + version
	weatherQry  = "%s/weather/?q=%s&units=%s"
	forecastQry = "%s/forecast/?q=%s&units=%s"
)

func Current(q string, units string) (*CurrentWeather, error) {
	weather := CurrentWeather{}
	url := fmt.Sprintf(weatherQry, baseUrl, q, units)
	err := request(url, &weather)
	return &weather, err
}

func Forecast(q string, units string) (*ForecastWather, error) {
	forecast := ForecastWather{}
	url := fmt.Sprintf(forecastQry, baseUrl, q, units)
	err := request(url, &forecast)
	return &forecast, err
}

func request(url string, v interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(v)
}
