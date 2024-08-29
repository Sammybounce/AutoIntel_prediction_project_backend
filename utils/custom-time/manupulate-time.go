package customTime

import (
	"time"
)

func AddMinutesToCurrentTime(mins time.Duration) time.Time {
	currentTime := time.Now()
	return currentTime.Add(mins * time.Minute)
}

func RemoveMinutesFromCurrentTime(mins time.Duration) time.Time {
	currentTime := time.Now()
	return currentTime.Add(-mins * time.Minute)
}

func AddHoursToCurrentTime(hours time.Duration) time.Time {
	currentTime := time.Now()
	return currentTime.Add(hours * time.Hour)
}

func RemoveHoursFromCurrentTime(hours time.Duration) time.Time {
	currentTime := time.Now()
	return currentTime.Add(-hours * time.Hour)
}

func AddDayFromCurrentTime(days int) time.Time {
	currentTime := time.Now()
	return currentTime.AddDate(0, 0, days)
}

func RemoveDayFromCurrentTime(days int) time.Time {
	currentTime := time.Now()
	return currentTime.AddDate(0, 0, -days)
}
