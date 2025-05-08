package handlers

type DailyTemperatureResponse struct {
	Date            string `json:"date"`
	Min             int    `json:"min"`
	MinLowerBound   int    `json:"minLowerBound"`
	MingHigherBound int    `json:"minHigherBound"`
	Max             int    `json:"max"`
	MaxLowerBound   int    `json:"maxLowerBound"`
	MaxHigherBound  int    `json:"maxHigherBound"`
}

type TemperatureSummaryResponse struct {
	Region       string                     `json:"region"`
	ForecastDays []DailyTemperatureResponse `json:"forecastDays"`
}
type VersionResponse struct {
	Version string `json:"version"`
}
