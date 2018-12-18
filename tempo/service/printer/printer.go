package printer

import (
	"fmt"
	"strings"
	"zimmer/tempo/model"
)

func PrintWeather(w model.Weather, unit string) {
	header := fmt.Sprintf("     %s, %s    (%.1f, %.1f)", w.Name, w.Sys.Country, w.Coord.Lon, w.Coord.Lat)
	unit = strings.ToUpper(unit)

	fmt.Printf("\n%s"+
		"\n%s"+
		"\n\n   Weather:    %s (%s)"+
		"\n\n   Temperature:    %.1f %s\n           Max:    %.1f %s\n           Min:    %.1f %s"+
		"\n\n   Clouds:         %.1f %%\n   Visibility:     %.1f %%"+
		"\n\n",
		header, lineMaker(len(header), 4),
		w.Weather[0].Main, w.Weather[0].Description,
		tempConvert(w.Main.Temp, unit), unit,
		tempConvert(w.Main.TempMax, unit), unit,
		tempConvert(w.Main.TempMin, unit), unit,
		w.Clouds.All, w.Visibility/100,
	)

}

func tempConvert(temp float64, unit string) float64 {

	switch unit {
	case "C":
		return temp - 273.15
	case "F":
		return (9/5)*(temp-273.15) + 32
	case "K":
		return temp
	}

	return 0
}

func lineMaker(strSize int, additional int) string {
	line := " "

	for i := 0; i < strSize-1+additional; i++ {
		line += "-"
	}

	return line
}
