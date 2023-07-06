# go weather cli

A simple interactive terminal-based app in Go that fetches weather data using the OpenWeatherMap API.

## Running the App

To use this app, follow these steps:

1. Install Go on your machine.
2. Get your OpenWeatherMap API key by registering on their [website](https://openweathermap.org/).
3. Clone this repository: `git clone https://github.com/MOHYAZZZ/go-weather-cli.git`.
4. Go to the cloned directory: `cd go-weather-cli`.
5. Open `main.go` in a text editor, find the line `apiKey := "<Your-API-Key>"`, and replace `<Your-API-Key>` with your actual API key.
6. Run the app by typing `go run main.go` in your terminal.

## How to Use

1. Type a city's name and press the "Get Weather" button.
2. The weather info for that city will appear.
3. To exit, press the "Quit" button.


Note: Weather data is obtained from the [OpenWeatherMap API](https://openweathermap.org/api). The terminal UI is created using the [tview package](https://github.com/rivo/tview).
