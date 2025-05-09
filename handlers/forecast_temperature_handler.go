package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/hippo-an/hivebox/config"
)

func getForcastTemperature() (*TemperatureSummaryResponse, error) {
	serviceKey := config.Config.ForecastServiceKey
	nextSixTime, err := getNextSixTime()

	if err != nil {
		return nil, fmt.Errorf("invalid next six time: %w", err)
	}

	url := fmt.Sprintf("http://apis.data.go.kr/1360000/MidFcstInfoService/getMidTa?serviceKey=%s&pageNo=1&numOfRows=10&dataType=JSON&regId=11B10101&tmFc=%s", serviceKey, nextSixTime)
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch forecast temperature: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response forecastTemperature

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to json decode forecast temperature: %w", err)
	}

	if len(response.Response.Body.Items.Item) == 0 {
		return nil, errors.New("empty forecast temperature result")
	}

	result := mapToTemperatureSummary(response.Response.Body.Items.Item[0])

	return &result, nil
}

func GetForecastTemperatureHandler(w http.ResponseWriter, r *http.Request) {
	result, err := getForcastTemperature()

	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get forecast temperature: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func getNextSixTime() (string, error) {
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		return "", fmt.Errorf("getNextSixTime - err from load location: %w", err)
	}

	now := time.Now().In(loc)
	morning := time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, loc)
	evening := time.Date(now.Year(), now.Month(), now.Day(), 18, 0, 0, 0, loc)

	if now.After(morning) && now.Before(evening) {
		return morning.Format("200601021504"), nil
	}

	if now.Before(morning) {
		return evening.Add(-24 * time.Hour).Format("200601021504"), nil
	}

	if now.After(evening) {
		return evening.Format("200601021504"), nil
	}

	return "", errors.New("unsupported time")
}

type forecastTemperature struct {
	Response struct {
		Header struct {
			ResultCode string `json:"resultCode"`
			ResultMsg  string `json:"resultMsg"`
		} `json:"header"`
		Body struct {
			DataType string `json:"dataType"`
			Items    struct {
				Item []forecastTemperatureItem `json:"item"`
			} `json:"items"`
			PageNo     int `json:"pageNo"`
			NumOfRows  int `json:"numOfRows"`
			TotalCount int `json:"totalCount"`
		} `json:"body"`
	} `json:"response"`
}

type forecastTemperatureItem struct {
	RegID      string `json:"regId"`
	TaMin5     int    `json:"taMin5"`
	TaMin5Low  int    `json:"taMin5Low"`
	TaMin5High int    `json:"taMin5High"`
	TaMax5     int    `json:"taMax5"`
	TaMax5Low  int    `json:"taMax5Low"`
	TaMax5High int    `json:"taMax5High"`

	TaMin6     int `json:"taMin6"`
	TaMin6Low  int `json:"taMin6Low"`
	TaMin6High int `json:"taMin6High"`
	TaMax6     int `json:"taMax6"`
	TaMax6Low  int `json:"taMax6Low"`
	TaMax6High int `json:"taMax6High"`

	TaMin7     int `json:"taMin7"`
	TaMin7Low  int `json:"taMin7Low"`
	TaMin7High int `json:"taMin7High"`
	TaMax7     int `json:"taMax7"`
	TaMax7Low  int `json:"taMax7Low"`
	TaMax7High int `json:"taMax7High"`

	TaMin8     int `json:"taMin8"`
	TaMin8Low  int `json:"taMin8Low"`
	TaMin8High int `json:"taMin8High"`
	TaMax8     int `json:"taMax8"`
	TaMax8Low  int `json:"taMax8Low"`
	TaMax8High int `json:"taMax8High"`

	TaMin9     int `json:"taMin9"`
	TaMin9Low  int `json:"taMin9Low"`
	TaMin9High int `json:"taMin9High"`
	TaMax9     int `json:"taMax9"`
	TaMax9Low  int `json:"taMax9Low"`
	TaMax9High int `json:"taMax9High"`

	TaMin10     int `json:"taMin10"`
	TaMin10Low  int `json:"taMin10Low"`
	TaMin10High int `json:"taMin10High"`
	TaMax10     int `json:"taMax10"`
	TaMax10Low  int `json:"taMax10Low"`
	TaMax10High int `json:"taMax10High"`
}

func mapToTemperatureSummary(item forecastTemperatureItem) TemperatureSummaryResponse {
	loc, _ := time.LoadLocation("Asia/Seoul")

	now := time.Now().In(loc)

	return TemperatureSummaryResponse{
		Region: item.RegID,
		ForecastDays: []DailyTemperatureResponse{
			{Date: now.AddDate(0, 0, 5).Format("2006-01-02"), Min: item.TaMin5, MinLowerBound: item.TaMin5Low, MingHigherBound: item.TaMin5High, Max: item.TaMax5, MaxLowerBound: item.TaMax5Low, MaxHigherBound: item.TaMax5High},
			{Date: now.AddDate(0, 0, 6).Format("2006-01-02"), Min: item.TaMin6, MinLowerBound: item.TaMin6Low, MingHigherBound: item.TaMin6High, Max: item.TaMax6, MaxLowerBound: item.TaMax6Low, MaxHigherBound: item.TaMax6High},
			{Date: now.AddDate(0, 0, 7).Format("2006-01-02"), Min: item.TaMin7, MinLowerBound: item.TaMin7Low, MingHigherBound: item.TaMin7High, Max: item.TaMax7, MaxLowerBound: item.TaMax7Low, MaxHigherBound: item.TaMax7High},
			{Date: now.AddDate(0, 0, 8).Format("2006-01-02"), Min: item.TaMin8, MinLowerBound: item.TaMin8Low, MingHigherBound: item.TaMin8High, Max: item.TaMax8, MaxLowerBound: item.TaMax8Low, MaxHigherBound: item.TaMax8High},
			{Date: now.AddDate(0, 0, 9).Format("2006-01-02"), Min: item.TaMin9, MinLowerBound: item.TaMin9Low, MingHigherBound: item.TaMin9High, Max: item.TaMax9, MaxLowerBound: item.TaMax9Low, MaxHigherBound: item.TaMax9High},
			{Date: now.AddDate(0, 0, 10).Format("2006-01-02"), Min: item.TaMin10, MinLowerBound: item.TaMin10Low, MingHigherBound: item.TaMin10High, Max: item.TaMax10, MaxLowerBound: item.TaMax10Low, MaxHigherBound: item.TaMax10High},
		},
	}
}
