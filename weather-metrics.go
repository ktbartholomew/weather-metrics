package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Metric is a single key/value pair
type Metric struct {
	Name  string
	Value string
}

type weatherReport struct {
	Altimeter     string
	Dewpoint      string
	Temperature   string
	WindDirection string `json:"Wind-Direction"`
	WindSpeed     string `json:"Wind-Speed"`
}

const up string = "1"

var weatherStation string
var requestCount int

func main() {
	var port string
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	} else {
		port = ":8080"
	}

	if os.Getenv("STATION") != "" {
		weatherStation = os.Getenv("STATION")
	} else {
		panic("Environment variable STATION not provided")
	}

	http.HandleFunc("/metrics", func(res http.ResponseWriter, req *http.Request) {
		requestCount++

		io.WriteString(res, formatMetrics(fetchMetrics()))
	})

	http.ListenAndServe(port, nil)
}

func fetchMetrics() []Metric {
	var metrics []Metric

	weather := getWeather()

	metrics = append(metrics, Metric{Name: "weather_metrics_up", Value: up})
	metrics = append(metrics, Metric{Name: "weather_metrics_request_count", Value: fmt.Sprint(requestCount)})
	metrics = append(metrics, Metric{Name: "weather_metrics_air_pressure", Value: weather.Altimeter})
	metrics = append(metrics, Metric{Name: "weather_metrics_temperature", Value: weather.Temperature})
	metrics = append(metrics, Metric{Name: "weather_metrics_wind_direction", Value: weather.WindDirection})
	metrics = append(metrics, Metric{Name: "weather_metrics_wind_speed", Value: weather.WindSpeed})

	return metrics
}

func formatMetrics(metrics []Metric) string {
	var metricsString string

	for i := 0; i < len(metrics); i++ {
		metricsString += fmt.Sprintf("%s %s\n", metrics[i].Name, metrics[i].Value)
	}

	return metricsString
}

func getWeather() weatherReport {
	weatherURL := fmt.Sprintf("https://avwx.rest/api/metar/%s", weatherStation)

	client := &http.Client{Timeout: time.Second * 10}

	res, err := client.Get(weatherURL)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	weather := weatherReport{}
	json.Unmarshal(buf, &weather)
  if weather.Temperature[0] == 77 {
    weather.Temperature = "-" + weather.Temperature[1:]
  }

	return weather
}
