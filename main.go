package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"tempo/args"
	"tempo/service/api"
	"tempo/service/printer"
)

func main() {

	var city string
	var country string
	var unit string
	var mode string

	args := args.Map(os.Args)
	processArgs(args, &city, &country, &unit, &mode)

	client := http.DefaultClient

	if mode == "weather" {

		w, err := api.DownloadWeather(client, api.DefaultUrlWeather(city, country))
		checkErr(err)

		printer.PrintWeather(w, unit)

	} else if mode == "forecast" {

		w, err := api.DownloadForecast(client, api.DefaultUrlForecast(city, country))
		checkErr(err)

		printer.PrintForecast(w, unit)

	} else {
		log.Fatal("Idk, man")
	}

}

func processArgs(args map[string]string, city *string, country *string, unit *string, mode *string) {

	if os.Args[1] == "forecast" {

		*mode = strings.ToLower("forecast")

		switch len(os.Args) {
		case 2:
			log.Fatal("Not enough arguments!")
		case 3:
			*city = os.Args[2]
			*country = "zzz"
			*unit = "C"
		case 4:
			*city = os.Args[2]
			*country = os.Args[3]
			*unit = "C"
		case 5:
			*city = os.Args[2]
			*country = os.Args[3]
			*unit = strings.ToUpper(os.Args[4])
		}

	} else if args[""] != "" {

		switch len(os.Args) {
		case 1:
			log.Fatal("Not enough arguments!")
		case 2:
			*city = os.Args[1]
			*country = "zzz"
			*unit = "C"
		case 3:
			*city = os.Args[1]
			*country = os.Args[2]
			*unit = "C"
		case 4:
			*city = os.Args[1]
			*country = os.Args[2]
			*unit = strings.ToUpper(os.Args[3])
		}
		*mode = "weather"

	} else {

		if args["t"] != "" {
			*city = args["t"]
		} else {
			log.Fatal("Not enough arguments!")
		}
		if args["c"] != "" {
			*country = args["c"]
		} else {
			*country = "zzz"
		}
		if args["u"] != "" {
			*unit = args["u"]
		} else {
			*unit = "C"
		}
		*mode = "weather"

	}
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
