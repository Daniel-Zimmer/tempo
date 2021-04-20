package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"tempo/model"
)

type ApiUrl struct {
	Base    string
	City    string
	Country string
	Appid   string
}

func DownloadWeather(client *http.Client, urlStruct ApiUrl) (model.Weather, error) {
	url := parseUrl(urlStruct)

	resp, err := client.Get(url)
	checkErr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	var w model.Weather

	if resp.StatusCode == 404 {
		return w, errors.New("City not found!")
	}

	err = json.Unmarshal(body, &w)
	checkErr(err)

	return w, nil
}

func DownloadForecast(client *http.Client, urlStruct ApiUrl) (model.Forecast, error) {
	url := parseUrl(urlStruct)

	resp, err := client.Get(url)
	checkErr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	var w model.Forecast

	if resp.StatusCode == 404 {
		return w, errors.New("City not found!")
	}

	err = json.Unmarshal(body, &w)
	checkErr(err)

	return w, nil
}

func DefaultUrlWeather(city string, country string) (url ApiUrl) {
	url.Base = "https://api.openweathermap.org/data/2.5/weather"
	url.City = formatCity(city)
	url.Country = strings.ToUpper(country)
	url.Appid = "e72d6671bfd19d515972372ca82287ef"

	return
}

func DefaultUrlForecast(city string, country string) (url ApiUrl) {
	url.Base = "https://api.openweathermap.org/data/2.5/forecast"
	url.City = formatCity(city)
	url.Country = strings.ToUpper(country)
	url.Appid = "e72d6671bfd19d515972372ca82287ef"

	return
}

func parseUrl(url ApiUrl) string {
	return url.Base + "?q=" + formatCity(url.City) + "," + url.Country +
		"&APPID=" + url.Appid
}

func formatCity(city string) string {
	newCity := ""

	for _, char := range city {
		if char == ' ' {
			newCity += "%20"
		} else {
			newCity += string(char)
		}
	}

	return newCity
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
