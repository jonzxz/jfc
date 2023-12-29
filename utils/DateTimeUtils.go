package utils

import (
	"strconv"
	"time"
)

func GetStartEndEpochFromMonth(month string) map[string]int64 {
	now := time.Now()
	currentYear, _, _ := now.Date()
	monthInt, _ := strconv.Atoi(month)
	monthTime := time.Month(monthInt)
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, monthTime, 1, 0, 0, 0, 0, currentLocation)
	// dirty add to make 23:59:59
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1).UTC().Add(time.Hour * 23).Add(time.Minute * 59).Add(time.Second * 59)

	startUnix := firstOfMonth.Unix()
	endUnix := lastOfMonth.Unix()

	epochs := make(map[string]int64)

	epochs["start"] = startUnix
	epochs["end"] = endUnix

	return epochs

}

func GetLastEpochOfCurrentMonthFromEpoch(epoch int64) int64 {
	eDate := time.Unix(epoch, 0)
	firstOfMonth := time.Date(eDate.Year(), eDate.Month(), 1, 0, 0, 0, 0, eDate.Location())
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1).UTC().Add(time.Hour * 23).Add(time.Minute * 59).Add(time.Second * 59)

	return lastOfMonth.Unix()
}
