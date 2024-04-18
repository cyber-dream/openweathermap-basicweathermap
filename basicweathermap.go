package openweathermap

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

type BasicWeatherLayer string

const (
	CloudsLayer        BasicWeatherLayer = "clouds_new"
	PrecipitationLayer BasicWeatherLayer = "precipitation_new"
	PressureLayer      BasicWeatherLayer = "pressure_new"
	WindLayer          BasicWeatherLayer = "wind_new"
	TemperatureLayer   BasicWeatherLayer = "temp_new"
)

type BasicWeatherMap struct {
	Key     string
	baseURL string
	*Settings
}

// NewBasicWeatherMap returns a new BasicWeatherMap pointer with the supplied arguments.
func NewBasicWeatherMap(key string) *BasicWeatherMap {
	settings := NewSettings()

	wmPtr := BasicWeatherMap{
		baseURL:  basicWeatherMapURL,
		Settings: settings,
		Key:      key,
	}

	return &wmPtr
}

// Precipitation will provide a tile with data.
func (wm *BasicWeatherMap) Precipitation(x int, y int, zoom int) (*bytes.Buffer, error) {
	url := fmt.Sprintf(wm.baseURL, PrecipitationLayer, zoom, x, y, wm.Key)
	response, err := wm.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("map get error status= %s", response.Status)
	}

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return buf, nil
}
