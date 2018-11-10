package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
)

func env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loding .env file")
	}
}

func main() {
	env_load()

	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/", func(writer http.ResponseWriter, request *http.Request) {
		city := strings.SplitN(request.URL.Path, "/", 3)[2]

		data, err := query(city)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(writer).Encode(data)
	})

	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

func query(city string) (weatherData, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city +
		"&APPID=" + os.Getenv("OPEN_WEATHER_MAP_KEY"))
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var data weatherData

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return weatherData{}, err
	}

	return data, nil
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}
