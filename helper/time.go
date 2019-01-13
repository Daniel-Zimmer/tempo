package helper

import (
	"fmt"
	"time"
)

func TimeDiff(dt int64) (str string) {
	t := time.Unix(dt, 0)
	timePassed := time.Since(t)
	minutes := timePassed.Minutes()

	if minutes < 800 {
		str = fmt.Sprintf("%.0f minutes ago", minutes)
	} else {
		str = "Way too long ago"
	}

	return
}

func FormatDate(t time.Time) string {
	wd := abbreviateWeekday(t.Weekday())

	d := fmt.Sprint(t.Day())
	if t.Day() < 10 {
		d = "0" + d
	}

	return fmt.Sprintf("%s %s (%s)",
		d, abbreviateMonth(t.Month()), wd,
	)
}

func ParseHour(t time.Time) string {
	hour := t.Hour()

	if hour < 10 {
		return fmt.Sprintf("0%dh", hour)
	} else {
		return fmt.Sprintf("%dh", hour)
	}
}

func ParseTime(dt int64) time.Time {
	return time.Unix(dt, 0)
}

func abbreviateMonth(month time.Month) string {
	return month.String()[:3]
}

func abbreviateWeekday(weekday time.Weekday) string {
	return weekday.String()[:3]
}
