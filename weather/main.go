package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	var q string
	fmt.Scanln(&q)
	if len(os.Args) >= 2 {
		q = os.Args[1]
	}
	res, err := http.Get(
		"http://api.weatherapi.com/v1/forecast.json?key=be44aceab453650242001&q=" + q + "&days=1&aqi=no&alerts=no",
	)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour
	fmt.Printf(
		"%s, %s: %.0fC, %s\n",
		location.Name,
		location.Country,
		current.Tempc,
		current.Condition.Text,
	)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()) {
			continue
		}
		message := fmt.Sprintf(
			"%s - %.0fC, %.0f%%, %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)
		fmt.Print(message)
	}
}
