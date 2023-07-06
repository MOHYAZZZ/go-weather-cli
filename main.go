package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rivo/tview"
)

type WeatherResponse struct {
	Name    string    `json:"name"`
	Main    Main      `json:"main"`
	Weather []Weather `json:"weather"`
}

type Main struct {
	Temperature float64 `json:"temp"`
}

type Weather struct {
	Description string `json:"description"`
}

func fetchWeatherData(city string) (WeatherResponse, error) {
	apiKey := "<Your-API-Key>" // Replace with your OpenWeatherMap API key
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s", city, apiKey)

	response, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return WeatherResponse{}, fmt.Errorf("failed to fetch weather data: HTTP status %d", response.StatusCode)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return WeatherResponse{}, err
	}

	var weatherResponse WeatherResponse
	err = json.Unmarshal(data, &weatherResponse)
	if err != nil {
		return WeatherResponse{}, err
	}

	return weatherResponse, nil
}

func main() {
	app := tview.NewApplication()

	form := tview.NewForm()
	form.AddInputField("City", "", 20, nil, nil)

	textView := tview.NewTextView()

	form.AddButton("Get Weather", func() {
		city := form.GetFormItem(0).(*tview.InputField).GetText()

		weatherResponse, err := fetchWeatherData(city)
		if err != nil {
			fmt.Fprintf(textView, "Failed to fetch weather data: %v\n", err)
			return
		}

		fmt.Fprintf(textView, "City: %s\n", weatherResponse.Name)
		fmt.Fprintf(textView, "Temperature: %.2f°C\n", weatherResponse.Main.Temperature)
		fmt.Fprintf(textView, "Description: %s\n\n", weatherResponse.Weather[0].Description)
	})

	form.AddButton("Quit", func() {
		app.Stop()
	})

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(form, 3, 1, true).
		AddItem(textView, 0, 1, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
