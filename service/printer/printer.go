package printer

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
	"time"
	"zimmer/tempo/helper"
	"zimmer/tempo/model"
)

type data struct {
	date        [6][8]time.Time
	weather     [6][8]string
	temperature [6][8]string
	humidity    [6][8]string
}

func PrintWeather(w model.Weather, unit string) {
	header := fmt.Sprintf("     %s, %s    (%.1f, %.1f)", w.Name, w.Sys.Country, w.Coord.Lon, w.Coord.Lat)
	unit = strings.ToUpper(unit)

	fmt.Printf("\n%s"+
		"\n%s"+
		"\n\n   Weather:    %s %s"+
		"\n\n   Temperature:    %.1f %s\n           Max:    %.1f %s\n           Min:    %.1f %s"+
		"\n\n   Clouds:         %.1f %%\n   Humidity:       %.1f %%\n   Visibility:     %.1f %%"+
		"\n\n   Last Update:    %s"+
		"\n\n",
		header, lineMaker(len(header), 4),
		w.Weather[0].Main, descParse(w.Weather[0].Description),
		tempConvert(w.Main.Temp, unit), unit,
		tempConvert(w.Main.TempMax, unit), unit,
		tempConvert(w.Main.TempMin, unit), unit,
		w.Clouds.All, w.Main.Humidity, w.Visibility/100,
		helper.TimeDiff(w.Dt),
	)
}

func PrintForecast(w model.Forecast, unit string) {
	spacer := 13

	drawHeader(&w)

	var info data
	initInfo(&info, &w, unit)

	n := 6
	if info.weather[5][0] == "*" {
		n = 5
	}

	for j := 0; j < n; j++ {
		color.Set(color.FgYellow, color.Bold)
		fmt.Printf("\n %s:\n", parseHeadDate(info.date[j]))
		color.Unset()
		drawTime(spacer, &info.weather)
		draw(&info, spacer, j)
	}

	fmt.Println()

}

/*******************Drawers**************************************/

func drawHeader(w *model.Forecast) {
	header := fmt.Sprintf("%s, %s   (%.1f, %.1f)", w.City.Name, w.City.Country, w.City.Coord.Lon, w.City.Coord.Lat)

	color.Set(color.FgRed)
	fmt.Printf("\n%s%s", intToSpace(33), header)
	color.Set(color.FgHiGreen, color.Italic)
	fmt.Printf("\n%sCaption: W - Weather,  T - Temperature,  H - Humidity", intToSpace(22))
	color.Unset()
}

func initInfo(info *data, w *model.Forecast, unit string) {
	{
		j := 0
		k := 0
		lastDay := helper.ParseTime(w.List[0].Dt).Day()
		for i := 0; i < w.Cnt; i++ {
			if lastDay != helper.ParseTime(w.List[i].Dt).Day() {
				j++
				k = 0
			}

			info.date[j][k] = helper.ParseTime(w.List[i].Dt)
			info.weather[j][k] = w.List[i].Weather[0].Main
			info.temperature[j][k] = fmt.Sprintf("%.1f %s", tempConvert(w.List[i].Main.Temp, unit), unit)
			info.humidity[j][k] = fmt.Sprintf("%d%%", w.List[i].Main.Humidity)

			lastDay = helper.ParseTime(w.List[i].Dt).Day()
			k++
		}
	}

	size := arraySize(info.weather[0])
	copy(info.date[0][8-size:], info.date[0][:size])
	copy(info.weather[0][8-size:], info.weather[0][:size])
	copy(info.temperature[0][8-size:], info.temperature[0][:size])
	copy(info.humidity[0][8-size:], info.humidity[0][:size])
	for i := 0; i < 8-size; i++ {
		info.date[0][i] = time.Time{}
		info.weather[0][i] = "*"
		info.temperature[0][i] = "*"
		info.humidity[0][i] = "*"
	}

	size = arraySize(info.weather[5])
	for i := size; i < 8; i++ {
		info.date[5][i] = time.Time{}
		info.weather[5][i] = "*"
		info.temperature[5][i] = "*"
		info.humidity[5][i] = "*"
	}

}

func draw(info *data, spacer int, j int) {
	desc := color.New(color.FgHiMagenta).SprintFunc()

	fmt.Print(desc("  W.  "))
	for i := 0; i < 8; i++ {
		fmt.Printf("%s%s", info.weather[j][i], spaceMaker(&info.weather[j][i], spacer))
	}
	fmt.Println()
	fmt.Print(desc("  T.  "))
	for i := 0; i < 8; i++ {
		fmt.Printf("%s%s", info.temperature[j][i], spaceMaker(&info.temperature[j][i], spacer))
	}
	fmt.Println()
	fmt.Print(desc("  H.  "))
	for i := 0; i < 8; i++ {
		fmt.Printf("%s%s", info.humidity[j][i], spaceMaker(&info.humidity[j][i], spacer))
	}
	fmt.Println()
}

func parseHeadDate(date [8]time.Time) string {
	if date[0].Weekday().String() != "" {
		return helper.FormatDate(date[0])
	}
	return helper.FormatDate(date[7])
}

func drawTime(spacer int, weather *[6][8]string) {
	fmt.Print("      ")

	grey := color.New(color.FgBlack).SprintFunc()
	color.Set(color.FgBlue)
	for i := 1; i < 10; i += 3 {
		fmt.Printf("0%dh%s", i, intToSpace(spacer-3))
	}
	for i := 10; i < 24; i += 3 {
		fmt.Printf("%dh%s", i, intToSpace(spacer-3))
	}
	color.Unset()

	fmt.Println()
}

/******************Utilities*************************************/

func arraySize(array [8]string) int {
	for i := 0; i < 8; i++ {
		if array[i] == "" {
			return i
		}
	}
	return 8
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

func descParse(desc string) string {
	str := strings.Split(desc, " ")

	if desc == "clear sky" {
		return ""
	}

	if len(str) == 1 {
		return ""
	} else if len(str) == 2 {
		return "(" + str[0] + ")"
	} else if len(str) >= 3 {
		if str[0] == "thunderstorm" {
			return "(" + strings.Join(str[1:], " ") + ")"
		} else {
			return "(" + strings.Join(str[:len(str)-1], " ") + ")"
		}
	}
	return "(" + desc + ")"
}

func lineMaker(strSize int, additional int) string {
	line := " "

	for i := 0; i < strSize-1+additional; i++ {
		line += "-"
	}

	return line
}

func spaceMaker(str *string, reqSize int) string {
	line := ""

	for i := 0; i < reqSize-len(*str); i++ {
		line += " "
	}

	return line
}

func intToSpace(n int) string {
	str := ""
	for i := 0; i < n; i++ {
		str += " "
	}
	return str
}
