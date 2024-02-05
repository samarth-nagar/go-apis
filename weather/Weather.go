package main

type Location struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	State   string `json:"state"`
	Street  string `json:"street"`
}
type (
	Current struct{}
	Weather struct {
		Location struct {
			Name    string `json:"name"`
			Country string `json:"country"`
		} `json:"location"`
		Current struct {
			Tempc     float64 `json:"temp_c"`
			Condition struct {
				Text string `json:"text"`
			}
		} `json:"current"`
		Forecast struct {
			Forecastday []struct {
				Hour []struct {
					TimeEpoch int64   `json:"time_epoch"`
					TempC     float64 `json:"temp_c"`
					Condition struct {
						Text string `json:"text"`
					}
					ChanceOfRain float64 `json:"chance_of_rain"`
				} `json:"hour"`
			} `json:"forecastday"`
		} `json:"forecast"`
	}
)
