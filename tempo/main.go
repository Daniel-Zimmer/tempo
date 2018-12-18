package main

import (
	"log"
	"net/http"
	"os"
	"zimmer/args"
	"zimmer/tempo/service/api"
	"zimmer/tempo/service/printer"
)

func main() {

	var city string
	var country string
	var unit string

	args := args.Map(os.Args)
	processArgs(args, &city, &country, &unit)

	client := http.DefaultClient

	w := api.Download(client, api.DefaultUrl(city, country))

	printer.PrintWeather(w, unit)
	
}

func processArgs(args map[string]string, city *string, country *string, unit *string) {
	if args[""] != "" {

		switch len(os.Args) {
		case 1:
			*city = "Rio de Janeiro"
			*country = "BR"
			*unit = "C"
		case 2:
			log.Fatal("Not enough arguments")
		case 3:
			*city = os.Args[1]
			*country = os.Args[2]
			*unit = "C"
		case 4:
			*city = os.Args[1]
			*country = os.Args[2]
			*unit = os.Args[3]
		}

	} else {

		if args["t"] != "" {
			*city = args["t"]
		} else {
			*city = "Rio de Janeiro"
		}
		if args["c"] != "" {
			*country = args["c"]
		} else {
			*country = "BR"
		}
		if args["u"] != "" {
			*unit = args["u"]
		} else {
			*unit = "C"
		}

	}
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}