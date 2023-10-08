package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/Kozzen890/assignment3-016/models"
)

func UpdateData() {
	for {
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1
		
		if err := RequestAPI(water, wind); err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	
		water_Data := getWeatherStatus(water, 6, 8)
		wind_Data := getWeatherStatus(wind, 7, 15)
	
		ResultData := models.WeatherData{
			Water: water,
			Wind:  wind,
		}
		
		logWeatherData(ResultData, water_Data, wind_Data)
		
		// update data tiap 15 detik
		time.Sleep(15 * time.Second)
	}

}

func logWeatherData(data models.WeatherData, waterStatus, windStatus string) {
	logJSON, _ := json.MarshalIndent(data, "", "  ")
	fmt.Printf("%s\nWater Status: %s\nWind Status: %s\n", string(logJSON), waterStatus, windStatus)
}

func RequestAPI(result_water, result_wind int) error {
	dataUpToDate := models.WeatherData{
		Water: result_water,
		Wind:  result_wind,
	}
	apiLink := "http://localhost:8080/apiUpdate"
	payload, _ := json.Marshal(dataUpToDate)
	
	req, err := http.NewRequest("PUT", apiLink, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error:", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("API returned non-OK status:", resp.Status)
	}

	return nil
}

func getWeatherStatus (data, minRange, maxRange int) string{
	var result string
	if data < minRange {
		result = "Aman"
	} else if data >= minRange && data <= maxRange {
		result = "Siaga"
	} else {
		result = "Bahaya"
	}
	return result
}