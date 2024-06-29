package utils

import (
	"strconv"
	"time"
)

func GetNowUTC() time.Time {
	loc, _ := time.LoadLocation("UTC")
	return time.Now().In(loc)
}

// 20240625
func GetYearMonthDay(t time.Time) int {
	date, _ := strconv.Atoi(t.Format("20060102"))
	return date
}

// 202447
func GetYearWeek(t time.Time) int {
	year, week := t.ISOWeek()
	date := year*100 + week
	return date
}
