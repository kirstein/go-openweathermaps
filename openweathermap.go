package openweathermap

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	version = "2.5"
	baseUrl = "http://api.openweathermap.org/data/" + version
)

func Weather(q string, units string) (*Current, error) {
	weather := Current{}
	url := fmt.Sprintf("%s/weather/?q=%s&units=%s", baseUrl, q, units)
	err := request(url, &weather)
	return &weather, err
}

func request(url string, v interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(v)
}
